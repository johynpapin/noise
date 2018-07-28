package things

func init() {
	registerThing("Float64Gen", NewFloat64Gen)
}

func NewFloat64Gen(name string) *Thing {
	t := newThing(name)

	t.Settings.add(t, "value", 0, 0, 127)

	t.Outputs.add(t, "output", NewFloat64GenOutput(t))

	return t
}

// Float64GenOutput

type Float64GenOutput struct {
	t *Thing
}

func NewFloat64GenOutput(t *Thing) *Float64GenOutput {
	return &Float64GenOutput{
		t: t,
	}
}

func (o Float64GenOutput) Next(sample int, entryPoint int) float64 {
	return o.t.Settings.Get("value").Value
}
