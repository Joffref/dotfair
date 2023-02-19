package provider

import (
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"github.com/dotfair-opensource/dotfair/pkg/provider/aws"
)

type IaCProvider struct {
	Name   string                 `json:"name" yaml:"name"`
	Probes map[string]probe.Probe `json:"probes" yaml:"probes"`
}

var (
	Providers = map[string]*IaCProvider{
		"aws": {
			Name:   "aws",
			Probes: aws.Probes,
		},
	}
)
