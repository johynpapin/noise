package things

import (
	"math"
)

func init() {
	registerThing("filter", NewFilter)
}

func NewFilter(name string) *Thing {
	t := newThing(name)

	t.Inputs.add(t, "signal", 0)
	t.Inputs.add(t, "cutoff", 0.99)
	t.Inputs.add(t, "resonance", 1)

	t.Outputs.add(t, "lowpass", NewFilterLowPass(t))
	t.Outputs.add(t, "highpass", NewFilterHighPass(t))
	t.Outputs.add(t, "bandpass", NewFilterBandPass(t))

	return t
}

// FilterLowPass

type FilterLowPass struct {
	t     *Thing
	state *State
}

func NewFilterLowPass(t *Thing) *FilterLowPass {
	return &FilterLowPass{
		t:     t,
		state: newState(),
	}
}

func (o *FilterLowPass) Next(sample int, entryPoint int) float64 {
	var (
		signal    = o.t.Inputs.Read("signal", sample, entryPoint)
		cutoff    = o.t.Inputs.Read("cutoff", sample, entryPoint)
		resonance = o.t.Inputs.Read("resonance", sample, entryPoint)

		lastCutoff = o.state.GetFloat64("lastCutoff", sample, entryPoint)
		g          = o.state.GetFloat64("g", sample, entryPoint)
		s1         = o.state.GetFloat64("s1", sample, entryPoint)
		s2         = o.state.GetFloat64("s2", sample, entryPoint)
	)

	var hp, bp, lp float64

	cutoff = math.Abs(cutoff)
	if cutoff != lastCutoff {
		g = math.Tan(cutoff)
		lastCutoff = cutoff
	}

	r := 1 / math.Max(resonance, 1)
	h := 1 / (1 + r*g + g*g)

	for j := 0; j < 4; j++ {
		hp = h * (signal - r*s1 - g*s1 - s2)
		bp = g*hp + s1
		lp = g*bp + s2

		s1 = g*hp + bp
		s2 = g*bp + lp
	}

	o.state.Set("lastCutoff", lastCutoff)
	o.state.Set("g", g)
	o.state.Set("s1", s1)
	o.state.Set("s2", s2)

	return lp
}

// FilterHighPass

type FilterHighPass struct {
	t     *Thing
	state *State
}

func NewFilterHighPass(t *Thing) *FilterHighPass {
	return &FilterHighPass{
		t:     t,
		state: newState(),
	}
}

func (o *FilterHighPass) Next(sample int, entryPoint int) float64 {
	var (
		signal    = o.t.Inputs.Read("signal", sample, entryPoint)
		cutoff    = o.t.Inputs.Read("cutoff", sample, entryPoint)
		resonance = o.t.Inputs.Read("resonance", sample, entryPoint)

		lastCutoff = o.state.GetFloat64("lastCutoff", sample, entryPoint)
		g          = o.state.GetFloat64("g", sample, entryPoint)
		s1         = o.state.GetFloat64("s1", sample, entryPoint)
		s2         = o.state.GetFloat64("s2", sample, entryPoint)
	)

	var hp, bp, lp float64

	cutoff = math.Abs(cutoff)
	if cutoff != lastCutoff {
		g = math.Tan(cutoff)
		lastCutoff = cutoff
	}

	r := 1 / math.Max(resonance, 1)
	h := 1 / (1 + r*g + g*g)

	for j := 0; j < 4; j++ {
		hp = h * (signal - r*s1 - g*s1 - s2)
		bp = g*hp + s1
		lp = g*bp + s2

		s1 = g*hp + bp
		s2 = g*bp + lp
	}

	o.state.Set("lastCutoff", lastCutoff)
	o.state.Set("g", g)
	o.state.Set("s1", s1)
	o.state.Set("s2", s2)

	return hp
}

// FilterBandPass

type FilterBandPass struct {
	t     *Thing
	state *State
}

func NewFilterBandPass(t *Thing) *FilterBandPass {
	return &FilterBandPass{
		t:     t,
		state: newState(),
	}
}

func (o *FilterBandPass) Next(sample int, entryPoint int) float64 {
	var (
		signal    = o.t.Inputs.Read("signal", sample, entryPoint)
		cutoff    = o.t.Inputs.Read("cutoff", sample, entryPoint)
		resonance = o.t.Inputs.Read("resonance", sample, entryPoint)

		lastCutoff = o.state.GetFloat64("lastCutoff", sample, entryPoint)
		g          = o.state.GetFloat64("g", sample, entryPoint)
		s1         = o.state.GetFloat64("s1", sample, entryPoint)
		s2         = o.state.GetFloat64("s2", sample, entryPoint)
	)

	var hp, bp, lp float64

	cutoff = math.Abs(cutoff)
	if cutoff != lastCutoff {
		g = math.Tan(cutoff)
		lastCutoff = cutoff
	}

	r := 1 / math.Max(resonance, 1)
	h := 1 / (1 + r*g + g*g)

	for j := 0; j < 4; j++ {
		hp = h * (signal - r*s1 - g*s1 - s2)
		bp = g*hp + s1
		lp = g*bp + s2

		s1 = g*hp + bp
		s2 = g*bp + lp
	}

	o.state.Set("lastCutoff", lastCutoff)
	o.state.Set("g", g)
	o.state.Set("s1", s1)
	o.state.Set("s2", s2)

	return bp
}
