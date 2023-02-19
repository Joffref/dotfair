package probe

import (
	"context"
	tfjson "github.com/hashicorp/terraform-json"
)

type Probe func(ctx context.Context, resource *tfjson.ResourceChange) (Metrics, error)
