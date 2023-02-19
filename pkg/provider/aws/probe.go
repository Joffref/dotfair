package aws

import (
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"github.com/dotfair-opensource/dotfair/pkg/provider/aws/ressources"
)

var (
	Probes = map[string]probe.Probe{
		"aws_instance": ressources.InstanceProbe,
	}
)
