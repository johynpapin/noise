package things

import (
	"math"
)

func init() {
	registerThing("Overload", NewOverload)
}

func NewOverload(name string) *Thing {
	t := newThing(name)

	t.Inputs.add(t, "signal", 0)
	t.Inputs.add(t, "gain", 1)

	t.Outputs.add(t, "output", NewOverloadOutput(t))

	return t
}

// OverloadOutput

type OverloadOutput struct {
	t *Thing
}

func NewOverloadOutput(t *Thing) *OverloadOutput {
	return &OverloadOutput{
		t: t,
	}
}

func (o OverloadOutput) Next(sample int, entryPoint int) float64 {
	var (
		signal = o.t.Inputs.Read("signal", sample, entryPoint)
		gain   = o.t.Inputs.Read("gain", sample, entryPoint)
	)

	return math.Copysign(1, signal*gain) * (1 - math.Exp(-math.Abs(signal*gain)))
}
