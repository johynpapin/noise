package things

import (
	"github.com/johynpapin/noise/pkg/constants"
	"math"
)

func init() {
	registerThing("PolyBLEPGen", NewPolyBLEPGen)
}

type PolyBLEPGen struct {
	*IOS
}

func NewPolyBLEPGen() Thing {
	t := &PolyBLEPGen{
		NewIOS(),
	}

	t.IOS.inputs.add(t, "freq")
	t.IOS.inputs.add(t, "amp")
	t.IOS.inputs.add(t, "offset")

	t.IOS.outputs.add(t, "sine", NewPolyBLEPSine(t))
	t.IOS.outputs.add(t, "saw", NewPolyBLEPSaw(t))

	return t
}

// PolyBLEPSine

type PolyBLEPSine struct {
	t     *PolyBLEPGen
	phase float64
}

func NewPolyBLEPSine(t *PolyBLEPGen) *PolyBLEPSine {
	return &PolyBLEPSine{
		t: t,
	}
}

func (o *PolyBLEPSine) Next() float64 {
	var (
		freq   = o.t.IOS.inputs.Read("freq")
		amp    = o.t.IOS.inputs.Read("amp")
		offset = o.t.IOS.inputs.Read("offset")
	)

	next := math.Sin(o.phase)

	o.phase = StepPhase(o.phase, freq)

	return amp*next + offset
}

// PolyBLEPSaw

type PolyBLEPSaw struct {
	t     *PolyBLEPGen
	phase float64
}

func NewPolyBLEPSaw(t *PolyBLEPGen) *PolyBLEPSaw {
	return &PolyBLEPSaw{
		t: t,
	}
}

func (o *PolyBLEPSaw) Next() float64 {
	var (
		freq   = o.t.IOS.inputs.Read("freq")
		amp    = o.t.IOS.inputs.Read("amp")
		offset = o.t.IOS.inputs.Read("offset")
	)

	next := (2 * o.phase / (2 * math.Pi)) - 1
	next -= PolyBLEP(o.phase/(2*math.Pi), freq)

	o.phase = StepPhase(o.phase, freq)

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
