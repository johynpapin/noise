package things

func init() {
	registerThing("Add", NewAdd)
}

func NewAdd(name string) *Thing {
	t := newThing(name)

	t.Inputs.add(t, "signal1", 0)
	t.Inputs.add(t, "signal2", 0)

	t.Outputs.add(t, "output", NewAddOutput(t))

	return t
}

// AddOutput

type AddOutput struct {
	t *Thing
}

func NewAddOutput(t *Thing) *AddOutput {
	return &AddOutput{
		t: t,
	}
}

func (o *AddOutput) Next(sample int, entryPoint int) float64 {
	var (
		signal1 = o.t.Inputs.Read("signal1", sample, entryPoint)
		signal2 = o.t.Inputs.Read("signal2", sample, entryPoint)
	)

	return signal1 + signal2
}
