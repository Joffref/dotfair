package terraform

import tfjson "github.com/hashicorp/terraform-json"

func ToResourceChange(resource *tfjson.StateResource) *tfjson.ResourceChange {
	return &tfjson.ResourceChange{
		Address:      resource.Address,
		Mode:         resource.Mode,
		Type:         resource.Type,
		Name:         resource.Name,
		Index:        resource.Index,
		ProviderName: resource.ProviderName,
		Change: &tfjson.Change{
			Actions: []tfjson.Action{
				tfjson.ActionCreate,
			},
			Before: nil,
			After:  resource.AttributeValues,
		},
	}
}
