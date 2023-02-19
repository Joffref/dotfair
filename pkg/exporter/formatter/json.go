package formatter

import (
	"encoding/json"
	"github.com/dotfair-opensource/dotfair/pkg/probe"
)

const (
	GWPDescription = "Global Warming Potential"
	GWPLink        = "https://en.wikipedia.org/wiki/Global_warming_potential"
	PEDescription  = "Power Efficiency"
	PELink         = "https://en.wikipedia.org/wiki/Primary_energy"
	ADPDescription = "Average Daily Power"
	ADPLink        = "https://www.designingbuildings.co.uk/wiki/Abiotic_depletion_potential"
	GlobalDesc     = "If you want to see more details about the metrics: https://boavizta.org/en/blog/empreinte-de-la-fabrication-d-un-serveur"
)

type JSONVerbose struct {
	JSONDocument
	PreviousMetrics []probe.Metrics `json:"previous_metrics" yaml:"previous_metrics"`
	NextMetrics     []probe.Metrics `json:"new_metrics" yaml:"new_metrics"`
}

type PreviousFootprint struct {
	PreviousGWP probe.Value `json:"previous_gwp" yaml:"previous_gwp"`
	PreviousPE  probe.Value `json:"previous_pe" yaml:"previous_pe"`
	PreviousADP probe.Value `json:"previous_adp" yaml:"previous_adp"`
}

type NextFootprint struct {
	NextGWP probe.Value `json:"new_gwp" yaml:"new_gwp"`
	NextPE  probe.Value `json:"new_pe" yaml:"new_pe"`
	NextADP probe.Value `json:"new_adp" yaml:"new_adp"`
}

type DiffFootprint struct {
	DiffGWP probe.Value `json:"diff_gwp" yaml:"diff_gwp"`
	DiffPE  probe.Value `json:"diff_pe" yaml:"diff_pe"`
	DiffADP probe.Value `json:"diff_adp" yaml:"diff_adp"`
}

type JSONDocument struct {
	PreviousFootprint `json:"previous_footprint" yaml:"previous_footprint"`
	NextFootprint     `json:"new_footprint" yaml:"new_footprint"`
	DiffFootprint     `json:"diff_footprint" yaml:"diff_footprint"`
	GWPDesc           string `json:"gwp_desc" yaml:"gwp_desc"`
	GWPLink           string `json:"gwp_link" yaml:"gwp_link"`
	PEDesc            string `json:"pe_desc" yaml:"pe_desc"`
	PELink            string `json:"pe_link" yaml:"pe_link"`
	ADPDesc           string `json:"adp_desc" yaml:"adp_desc"`
	ADPLink           string `json:"adp_link" yaml:"adp_link"`
	Info              string `json:"info" yaml:"info"`
}

func JSON(previous, new []probe.Metrics, verbose bool) ([]byte, error) {
	previousGWP, previousPE, previousADP := probe.Sum(previous...)
	newGWP, newPE, newADP := probe.Sum(new...)
	jsonOutput := JSONDocument{
		PreviousFootprint: PreviousFootprint{
			PreviousGWP: previousGWP,
			PreviousPE:  previousPE,
			PreviousADP: previousADP,
		},
		NextFootprint: NextFootprint{
			NextGWP: newGWP,
			NextPE:  newPE,
			NextADP: newADP,
		},
		DiffFootprint: DiffFootprint{
			DiffGWP: probe.Diff(previousGWP, newGWP),
			DiffPE:  probe.Diff(previousPE, newPE),
			DiffADP: probe.Diff(previousADP, newADP),
		},
		GWPDesc: GWPDescription,
		GWPLink: GWPLink,
		PEDesc:  PEDescription,
		PELink:  PELink,
		ADPDesc: ADPDescription,
		ADPLink: ADPLink,
		Info:    GlobalDesc,
	}
	if verbose {
		jsonOutputVerbose := JSONVerbose{
			JSONDocument:    jsonOutput,
			PreviousMetrics: previous,
			NextMetrics:     new,
		}
		marshal, err := json.Marshal(jsonOutputVerbose)
		if err != nil {
			return nil, err
		}
		return marshal, nil
	}
	marshal, err := json.Marshal(jsonOutput)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}
