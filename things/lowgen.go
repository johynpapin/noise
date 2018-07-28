package things

import (
	"math"
)

func init() {
	registerThing("LowGen", NewLowGen)
}

func NewLowGen(name string) *Thing {
	t := newThing(name)

	t.Inputs.add(t, "freq", 440)
	t.Inputs.add(t, "amp", 1)
	t.Inputs.add(t, "pw", 1)
	t.Inputs.add(t, "offset", 0)

	t.Outputs.add(t, "sine", NewLowGenSine(t))
	t.Outputs.add(t, "pulse", NewLowGenPulse(t))
	t.Outputs.add(t, "saw", NewLowGenSaw(t))

	return t
}

// LowGenSine

type LowGenSine struct {
	t     *Thing
	phase float64
}

func NewLowGenSine(t *Thing) *LowGenSine {
	return &LowGenSine{
		t: t,
	}
}

func (o LowGenSine) Next(sample int, entryPoint int) float64 {
	var (
		freq   = o.t.Inputs.Read("freq", sample, entryPoint)
		amp    = o.t.Inputs.Read("amp", sample, entryPoint)
		offset = o.t.Inputs.Read("offset", sample, entryPoint)
	)

	next := math.Sin(o.phase)

	o.phase = StepPhase(o.phase, freq)

	return amp*next + offset
}

// LowGenPulse

type LowGenPulse struct {
	t     *Thing
	phase float64
}

func NewLowGenPulse(t *Thing) *LowGenPulse {
	return &LowGenPulse{
		t: t,
	}
}

func (o LowGenPulse) Next(sample int, entryPoint int) float64 {
	var (
		freq   = o.t.Inputs.Read("freq", sample, entryPoint)
		amp    = o.t.Inputs.Read("amp", sample, entryPoint)
		pw     = o.t.Inputs.Read("pw", sample, entryPoint)
		offset = o.t.Inputs.Read("offset", sample, entryPoint)
		next   float64
	)

	if o.phase < math.Pi*pw {
		next = 1
	} else {
		next = -1
	}

	o.phase = StepPhase(o.phase, freq)

	return amp*next + offset
}

// LowGenSaw

type LowGenSaw struct {
	t     *Thing
	phase float64
}

func NewLowGenSaw(t *Thing) *LowGenSaw {
	return &LowGenSaw{
		t: t,
	}
}

func (o LowGenSaw) Next(sample int, entryPoint int) float64 {
	var (
		freq   = o.t.Inputs.Read("freq", sample, entryPoint)
		amp    = o.t.Inputs.Read("amp", sample, entryPoint)
		offset = o.t.Inputs.Read("offset", sample, entryPoint)
	)

	next := (2*o.phase)/(2*math.Pi) - 1

	o.phase = StepPhase(o.phase, freq)

	return amp*next + offset
}
