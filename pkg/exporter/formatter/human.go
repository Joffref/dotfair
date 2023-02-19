package formatter

import (
	"bytes"
	"github.com/dotfair-opensource/dotfair/pkg/probe"
	"text/template"
)

var minimalStdoutTemplate = `
Previous Global Warming Potential: {{.PreviousGWP.Manufacture}} {{.PreviousGWP.Unit}} (Manufacture) + {{.PreviousGWP.Use}} {{.PreviousGWP.Unit}} (Use) = {{add .PreviousGWP.Manufacture .PreviousGWP.Use}} {{.PreviousGWP.Unit}}
Previous Primary Energy: {{.PreviousPE.Manufacture}} {{.PreviousPE.Unit}} (Manufacture) + {{.PreviousPE.Use}} {{.PreviousPE.Unit}} (Use) = {{add .PreviousPE.Manufacture .PreviousPE.Use}} {{.PreviousPE.Unit}}
Previous Abiotic Depletion Potential: {{.PreviousADP.Manufacture}} {{.PreviousADP.Unit}} (Manufacture) + {{.PreviousADP.Use}} {{.PreviousADP.Unit}} (Use) = {{add .PreviousADP.Manufacture .PreviousADP.Use}} {{.PreviousADP.Unit}}
Next Global Warming Potential: {{.NextGWP.Manufacture}} {{.NextGWP.Unit}} (Manufacture) + {{.NextGWP.Use}} {{.NextGWP.Unit}} (Use) = {{add .NextGWP.Manufacture .NextGWP.Use}} {{.NextGWP.Unit}}
Next Primary Energy: {{ .NextPE.Manufacture}} {{.NextPE.Unit}} (Manufacture) + {{.NextPE.Use}} {{.NextPE.Unit}} (Use) = {{add .NextPE.Manufacture .NextPE.Use}} {{.NextPE.Unit}}
Next Abiotic Depletion Potential: {{ .NextADP.Manufacture}} {{.NextADP.Unit}} (Manufacture) + {{.NextADP.Use}} {{.NextADP.Unit}} (Use) = {{add .NextADP.Manufacture .NextADP.Use}} {{.NextADP.Unit}}
Diff Global Warming Potential: {{sub (add .NextGWP.Manufacture .NextGWP.Use) (add .PreviousGWP.Manufacture .PreviousGWP.Use)}} {{.NextGWP.Unit}}
Diff Primary Energy: {{sub (add .NextPE.Manufacture .NextPE.Use) (add .PreviousPE.Manufacture .PreviousPE.Use)}} {{.NextPE.Unit}}
Diff Abiotic Depletion Potential: {{sub (add .NextADP.Manufacture .NextADP.Use) (add .PreviousADP.Manufacture .PreviousADP.Use)}} {{.NextADP.Unit}}
{{ if .Verbose }}{{ range $i, $e := .PreviousMetrics }}{{ $e.Provider }} - {{ $e.Resource }}:
  - Global Warming Potential: {{ $e.GWP.Manufacture }} {{ $e.GWP.Unit }} (Manufacture) + {{ $e.GWP.Use }} {{ $e.GWP.Unit }} (Use) = {{ add $e.GWP.Manufacture $e.GWP.Use }} {{ $e.GWP.Unit }}
  - Primary Energy: {{ $e.PE.Manufacture }} {{ $e.PE.Unit }} (Manufacture) + {{ $e.PE.Use }} {{ $e.PE.Unit }} (Use) = {{ add $e.PE.Manufacture $e.PE.Use }} {{ $e.PE.Unit }}
  - Abiotic Depletion Potential: {{ $e.ADP.Manufacture }} {{ $e.ADP.Unit }} (Manufacture) + {{ $e.ADP.Use }} {{ $e.ADP.Unit }} (Use) = {{ add $e.ADP.Manufacture $e.ADP.Use }} {{ $e.ADP.Unit }}
{{ end }}{{ range $i, $e := .NextMetrics }}{{ $e.Provider }} - {{ $e.Resource }}:
  + Global Warming Potential: {{ $e.GWP.Manufacture }} {{ $e.GWP.Unit }} (Manufacture) + {{ $e.GWP.Use }} {{ $e.GWP.Unit }} (Use) = {{ add $e.GWP.Manufacture $e.GWP.Use }} {{ $e.GWP.Unit }}
  + Primary Energy: {{ $e.PE.Manufacture }} {{ $e.PE.Unit }} (Manufacture) + {{ $e.PE.Use }} {{ $e.PE.Unit }} (Use) = {{ add $e.PE.Manufacture $e.PE.Use }} {{ $e.PE.Unit }}
  + Abiotic Depletion Potential: {{ $e.ADP.Manufacture }} {{ $e.ADP.Unit }} (Manufacture) + {{ $e.ADP.Use }} {{ $e.ADP.Unit }} (Use) = {{ add $e.ADP.Manufacture $e.ADP.Use }} {{ $e.ADP.Unit }}
{{ end }}{{ end }}{{.Note}}
`

var note = `Note:
- GWP: Global Warming Potential (kg CO2 eq) evaluates the effect on global warming. - More information: https://en.wikipedia.org/wiki/Global_warming_potential
- PE: Primary Energy (MJ) evaluates energy resources consumption(renewable + non-renewable) - More information: https://en.wikipedia.org/wiki/Primary_energy
- ADP: Abiotic Depletion Potential (kgSbeq) assesses the use of fossil minerals and raw materials. - More information: https://www.designingbuildings.co.uk/wiki/Abiotic_depletion_potential
If you want to see more details about the metrics: https://boavizta.org/en/blog/empreinte-de-la-fabrication-d-un-serveur`

var funcMap = template.FuncMap{
	"add": add,
	"sub": sub,
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func HumanReadable(previous, next []probe.Metrics, verbose bool) ([]byte, error) {
	previousGWP, previousPE, previousADP := probe.Sum(previous...)
	nextGWP, nextPE, nextADP := probe.Sum(next...)
	tmpl := template.Must(template.New("human").Funcs(funcMap).Parse(minimalStdoutTemplate))
	var output bytes.Buffer
	err := tmpl.Execute(&output, struct {
		PreviousGWP, PreviousPE, PreviousADP, NextGWP, NextPE, NextADP probe.Value
		PreviousMetrics, NextMetrics                                   []probe.Metrics
		Note                                                           string
		Verbose                                                        bool
	}{
		PreviousGWP:     previousGWP,
		PreviousPE:      previousPE,
		PreviousADP:     previousADP,
		NextGWP:         nextGWP,
		NextPE:          nextPE,
		NextADP:         nextADP,
		PreviousMetrics: previous,
		NextMetrics:     next,
		Note:            note,
		Verbose:         verbose,
	})
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
