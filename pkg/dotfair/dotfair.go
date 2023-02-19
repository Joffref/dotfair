package dotfair

import (
	"context"
	"fmt"
	"github.com/dotfair-opensource/dotfair/pkg/exporter/formatter"
	"github.com/dotfair-opensource/dotfair/pkg/exporter/stdout"
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"github.com/dotfair-opensource/dotfair/pkg/provider"
	"github.com/dotfair-opensource/dotfair/pkg/terraform"
	"strings"
)

type Config struct {
	Folder       string `json:"folder" yaml:"folder"`               // Folder is the folder to scan containing terraform code
	Verbose      bool   `json:"verbose" yaml:"verbose"`             // Verbose is the flag to enable verbose output
	OutputFormat string `json:"output_format" yaml:"output_format"` // OutputFormat is the format of the output
}

type Runner struct {
	config *Config
}

func (c *Config) Validate() error {
	if c.Folder == "" {
		return fmt.Errorf("folder is required")
	}
	if c.OutputFormat == "" {
		c.OutputFormat = "human-readable"
	}
	if c.OutputFormat != "human-readable" && c.OutputFormat != "json" {
		return fmt.Errorf("output_format must be either human-readable or json")
	}
	return nil
}

func (c *Config) SetDefaults() {
	if c.Folder == "" {
		c.Folder = "."
	}
}

func NewRunner(config *Config) (*Runner, error) {
	config.SetDefaults()
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &Runner{
		config: config,
	}, nil
}

func (r *Runner) Run(ctx context.Context) error {
	var beforeChange, afterChange []probe.Metrics
	state, err := terraform.CurrentState(ctx, r.config.Folder)
	if err != nil {
		return err
	}
	for _, resouce := range state.Values.RootModule.Resources {
		providerName := strings.Split(resouce.ProviderName, "/")
		name := providerName[len(providerName)-1]
		provider := provider.Providers[name]
		if provider == nil {
			continue
		}
		p := provider.Probes[resouce.Type]
		if p != nil {
			metrics, err := p(ctx, terraform.ToResourceChange(resouce))
			if err != nil {
				return err
			}
			beforeChange = append(beforeChange, metrics)
		}
	}
	plan, err := terraform.Plan(ctx, r.config.Folder)
	if err != nil {
		return err
	}
	for _, resouce := range plan.ResourceChanges {
		providerName := strings.Split(resouce.ProviderName, "/")
		name := providerName[len(providerName)-1]
		provider := provider.Providers[name]
		if provider == nil {
			continue
		}
		p := provider.Probes[resouce.Type]
		if p != nil {
			metrics, err := p(ctx, resouce)
			if err != nil {
				return err
			}
			afterChange = append(afterChange, metrics)
		}
	}
	if r.config.OutputFormat == "json" {
		stdout := stdout.Exporter{
			Formatter: formatter.JSON,
			Before:    beforeChange,
			After:     afterChange,
		}
		if err := stdout.Export(r.config.Verbose); err != nil {
			return err
		}
		return nil
	}
	stdout := stdout.Exporter{
		Formatter: formatter.HumanReadable,
		Before:    beforeChange,
		After:     afterChange,
	}
	if err := stdout.Export(r.config.Verbose); err != nil {
		return err
	}
	return nil
}
