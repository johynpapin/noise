package things

import (
	"github.com/johynpapin/noise/constants"
	"math"
)

func init() {
	registerThing("ADSR", NewADSR)
}

func NewADSR(name string) *Thing {
	t := newThing(name)

	t.Inputs.add(t, "trigger", 0)

	t.Inputs.add(t, "attack", 0)
	t.Inputs.add(t, "decay", 0)
	t.Inputs.add(t, "release", 0)

	t.Inputs.add(t, "sustain", 1)
	t.Inputs.add(t, "ratioA", 0.3)
	t.Inputs.add(t, "ratioDR", 0.0001)

	t.Outputs.add(t, "output", NewADSROutput(t))

	return t
}

type ADSRState int

const (
	Idle ADSRState = iota
	Attack
	Decay
	Sustain
	Release
)

// ADSROutput

type ADSROutput struct {
	t     *Thing
	state *State
}

func NewADSROutput(t *Thing) *ADSROutput {
	return &ADSROutput{
		t:     t,
		state: newState(),
	}
}

func (o *ADSROutput) Next(sample int, entryPoint int) float64 {
	var (
		trigger = o.t.Inputs.Read("trigger", sample, entryPoint)
		attack  = o.t.Inputs.Read("attack", sample, entryPoint) * constants.SAMPLE_RATE
		decay   = o.t.Inputs.Read("decay", sample, entryPoint) * constants.SAMPLE_RATE
		release = o.t.Inputs.Read("release", sample, entryPoint) * constants.SAMPLE_RATE
		sustain = o.t.Inputs.Read("sustain", sample, entryPoint)
		ratioA  = o.t.Inputs.Read("ratioA", sample, entryPoint)
		ratioDR = o.t.Inputs.Read("ratioDR", sample, entryPoint)

		state      = o.state.GetADSRState("state", sample, entryPoint)
		lastOutput = o.state.GetFloat64("lastOutput", sample, entryPoint)
	)

	if ratioA < 0.000000001 {
		ratioA = 0.000000001
	}

	if ratioDR < 0.0000000001 {
		ratioDR = 0.0000000001
	}

	attackCoef := calcCoef(attack, ratioA)
	attackBase := (1 + ratioA) * (1 - attackCoef)

	decayCoef := calcCoef(decay, ratioDR)
	decayBase := (sustain - ratioDR) * (1 - decayCoef)

	releaseCoef := calcCoef(release, ratioDR)
	releaseBase := -ratioDR * (1 - releaseCoef)

	if trigger > 0 && state == Idle {
		state = Attack
	} else if trigger <= 0 && state != Idle {
		state = Release
	}

	switch state {
	case Idle:
		lastOutput = 0
	case Attack:
		lastOutput = attackBase + lastOutput*attackCoef

		if lastOutput >= 1 {
			lastOutput = 1
			state = Decay
		}
	case Decay:
		lastOutput = decayBase + lastOutput*decayCoef

		if lastOutput <= sustain {
			lastOutput = sustain
			state = Sustain
		}
	case Sustain:
	case Release:
		lastOutput = releaseBase + lastOutput*releaseCoef

		if lastOutput < 0 {
			lastOutput = 0
			state = Idle
		}
	}

	o.state.Set("state", state)
	o.state.Set("lastOutput", lastOutput)

	return lastOutput
}

func calcCoef(rate float64, ratio float64) float64 {
	if rate <= 0 {
		return 0
	} else {
		return math.Exp(-math.Log((ratio+1)/ratio) / rate)
	}
}
