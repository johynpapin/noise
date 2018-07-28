package things

import (
	"github.com/johynpapin/noise/constants"
	"math"
)

func init() {
	registerThing("PolyBLEPGen", NewPolyBLEPGen)
}

func NewPolyBLEPGen(name string) *Thing {
	t := newThing(name)

	t.Inputs.add(t, "freq", 440)
	t.Inputs.add(t, "amp", 1)
	t.Inputs.add(t, "offset", 0)
	t.Inputs.add(t, "fm", 0)
	t.Inputs.add(t, "pm", 0)

	t.Outputs.add(t, "sine", NewPolyBLEPSine(t))
	t.Outputs.add(t, "saw", NewPolyBLEPSaw(t))

	return t
}

// PolyBLEPSine

type PolyBLEPSine struct {
	t     *Thing
	state *State
}

func NewPolyBLEPSine(t *Thing) *PolyBLEPSine {
	return &PolyBLEPSine{
		t:     t,
		state: newState(),
	}
}

func (o *PolyBLEPSine) Next(sample int, entryPoint int) float64 {
	var (
		freq   = o.t.Inputs.Read("freq", sample, entryPoint)
		amp    = o.t.Inputs.Read("amp", sample, entryPoint)
		offset = o.t.Inputs.Read("offset", sample, entryPoint)
		fm     = o.t.Inputs.Read("fm", sample, entryPoint)
		pm     = o.t.Inputs.Read("pm", sample, entryPoint)

		phase = o.state.GetFloat64("phase", sample, entryPoint)
	)

	next := math.Sin(phase + pm)

	o.state.Set("phase", StepPhaseWithFm(phase, freq, fm))

	return amp*next + offset
}

// PolyBLEPSaw

type PolyBLEPSaw struct {
	t     *Thing
	state *State
}

func NewPolyBLEPSaw(t *Thing) *PolyBLEPSaw {
	return &PolyBLEPSaw{
		t:     t,
		state: newState(),
	}
}

func (o *PolyBLEPSaw) Next(sample int, entryPoint int) float64 {
	var (
		freq   = o.t.Inputs.Read("freq", sample, entryPoint)
		amp    = o.t.Inputs.Read("amp", sample, entryPoint)
		offset = o.t.Inputs.Read("offset", sample, entryPoint)
		fm     = o.t.Inputs.Read("fm", sample, entryPoint)
		pm     = o.t.Inputs.Read("pm", sample, entryPoint)

		phase = o.state.GetFloat64("phase", sample, entryPoint)
	)

	next := (2 * (phase + pm) / (2 * math.Pi)) - 1
	next -= PolyBLEP(phase/(2*math.Pi), freq)

	o.state.Set("phase", StepPhaseWithFm(phase, freq, fm))

	return amp*next + offset
}

// utils

func PolyBLEP(phase, freq float64) float64 {
	delta := freq / constants.SAMPLE_RATE

	if phase < delta {
		phase /= delta
		return phase + phase - phase*phase - 1
	} else if phase > 1-delta {
		phase = (phase - 1) / delta
		return phase + phase + phase*phase + 1
	}

	return 0
}
