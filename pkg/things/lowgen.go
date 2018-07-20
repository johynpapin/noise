package things

import (
	"math"
)

func init() {
	registerThing("LowGen", NewLowGen)
}

type LowGen struct {
	*IOS
}

func NewLowGen() Thing {
	t := &LowGen{
		IOS: NewIOS(),
	}

	t.IOS.inputs.add(t, "freq")
	t.IOS.inputs.add(t, "amp")
	t.IOS.inputs.add(t, "pw")
	t.IOS.inputs.add(t, "offset")

	t.IOS.outputs.add(t, "sine", NewLowGenSine(t))
	t.IOS.outputs.add(t, "pulse", NewLowGenPulse(t))
	t.IOS.outputs.add(t, "saw", NewLowGenSaw(t))

	return t
}

// LowGenSine

type LowGenSine struct {
	t     *LowGen
	phase float64
}

func NewLowGenSine(t *LowGen) *LowGenSine {
	return &LowGenSine{
		t: t,
	}
}

func (o LowGenSine) Next() float64 {
	var (
		freq   = o.t.IOS.inputs.Read("freq")
		amp    = o.t.IOS.inputs.Read("amp")
		offset = o.t.IOS.inputs.Read("offset")
	)

	next := math.Sin(o.phase)

	o.phase = StepPhase(o.phase, freq)

	return amp*next + offset
}

// LowGenPulse

type LowGenPulse struct {
	t     *LowGen
	phase float64
}

func NewLowGenPulse(t *LowGen) *LowGenPulse {
	return &LowGenPulse{
		t: t,
	}
}

func (o LowGenPulse) Next() float64 {
	var (
		freq   = o.t.IOS.inputs.Read("freq")
		amp    = o.t.IOS.inputs.Read("amp")
		pw     = o.t.IOS.inputs.Read("pw")
		offset = o.t.IOS.inputs.Read("offset")
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
	t     *LowGen
	phase float64
}

func NewLowGenSaw(t *LowGen) *LowGenSaw {
	return &LowGenSaw{
		t: t,
	}
}

func (o LowGenSaw) Next() float64 {
	var (
		freq   = o.t.IOS.inputs.Read("freq")
		amp    = o.t.IOS.inputs.Read("amp")
		offset = o.t.IOS.inputs.Read("offset")
	)

	next := (2*o.phase)/(2*math.Pi) - 1

	o.phase = StepPhase(o.phase, freq)

	return amp*next + offset
}
