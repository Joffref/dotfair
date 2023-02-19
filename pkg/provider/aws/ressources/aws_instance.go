package ressources

import (
	"context"
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"github.com/dotfair-opensource/dotfair/pkg/provider/aws/source"
	tfjson "github.com/hashicorp/terraform-json"
)

func InstanceProbe(ctx context.Context, resource *tfjson.ResourceChange) (probe.Metrics, error) {
	var instanceType string
	if resource.Change.Actions.Delete() {
		switch resource.Change.Before.(type) {
		case map[string]interface{}:
			changes := resource.Change.Before.(map[string]interface{})
			if changes["instance_type"] != nil {
				instanceType = changes["instance_type"].(string)
				return deleteInstance(ctx, instanceType)
			}
		}
	} else {
		switch resource.Change.After.(type) {
		case map[string]interface{}:
			changes := resource.Change.After.(map[string]interface{})
			if changes["instance_type"] != nil {
				instanceType = changes["instance_type"].(string)
				return createInstance(ctx, instanceType)
			}
		}
	}
	return probe.Metrics{}, nil
}

func deleteInstance(ctx context.Context, instanceType string) (probe.Metrics, error) {
	client := source.NewBoaviztaApiClient("cloud", map[string]string{
		"provider":      source.ProviderName,
		"instance_type": instanceType,
		"verbose":       "false",
		"location":      "TOTAL",
	})
	resp, err := client.GetMetrics()
	if err != nil {
		return probe.Metrics{}, err
	}
	return probe.Metrics{
		Provider: source.ProviderName,
		Resource: "aws_instance",
		GWP: probe.Value{
			Unit:        resp.GWP.Unit,
			Manufacture: resp.GWP.Manufacture * -1,
			Use:         resp.GWP.Use * -1,
		},
		PE: probe.Value{
			Unit:        resp.PE.Unit,
			Manufacture: resp.PE.Manufacture * -1,
			Use:         resp.PE.Use * -1,
		},
		ADP: probe.Value{
			Unit:        resp.ADP.Unit,
			Manufacture: resp.ADP.Manufacture * -1,
			Use:         resp.ADP.Use * -1,
		},
	}, nil
}

func createInstance(ctx context.Context, instanceType string) (probe.Metrics, error) {
	client := source.NewBoaviztaApiClient("cloud", map[string]string{
		"provider":      source.ProviderName,
		"instance_type": instanceType,
		"verbose":       "false",
		"location":      "TOTAL",
	})
	resp, err := client.GetMetrics()
	if err != nil {
		return probe.Metrics{}, err
	}
	return probe.Metrics{
		Provider: source.ProviderName,
		Resource: "aws_instance",
		GWP:      resp.GWP,
		PE:       resp.PE,
		ADP:      resp.ADP,
	}, nil
}
