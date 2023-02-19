package probe

type Metrics struct {
	Provider string `json:"provider" yaml:"provider"`
	Resource string `json:"resource" yaml:"resource"`
	GWP      Value  `json:"gwp" yaml:"gwp"`
	PE       Value  `json:"pe" yaml:"pe"`
	ADP      Value  `json:"adp" yaml:"adp"`
}

type Value struct {
	Manufacture float64 `json:"manufacture" yaml:"manufacture"`
	Use         float64 `json:"use" yaml:"use"`
	Unit        string  `json:"unit" yaml:"unit"`
}

// Sum returns the sum of the values in the metrics. GWP, PE and ADP
func Sum(metrics ...Metrics) (Value, Value, Value) {
	var gwp, pe, adp Value
	for _, m := range metrics {
		gwp.Manufacture += m.GWP.Manufacture
		gwp.Use += m.GWP.Use
		pe.Manufacture += m.PE.Manufacture
		pe.Use += m.PE.Use
		adp.Manufacture += m.ADP.Manufacture
		adp.Use += m.ADP.Use
	}
	gwp.Unit = "kg CO2 eq"
	pe.Unit = "MJ"
	adp.Unit = "kgSbeq"
	return gwp, pe, adp
}

func Diff(before, after Value) Value {
	return Value{
		Manufacture: after.Manufacture - before.Manufacture,
		Use:         after.Use - before.Use,
		Unit:        before.Unit,
	}
}
