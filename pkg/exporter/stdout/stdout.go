package stdout

import (
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"os"
)

type Exporter struct {
	Formatter func(before, after []probe.Metrics, verbose bool) ([]byte, error)
	Before    []probe.Metrics
	After     []probe.Metrics
}

func (e *Exporter) Export(verbose bool) error {
	output, err := e.Formatter(e.Before, e.After, verbose)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(output)
	if err != nil {
		return err
	}
	return nil
}
