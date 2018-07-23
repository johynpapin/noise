package things

import (
	"github.com/johynpapin/noise/pkg/constants"
	"math"
)

func init() {
	registerThing("ADSR", NewADSR)
}

type ADSR struct {
	*IOS
}

func NewADSR() Thing {
	t := &ADSR{
		NewIOS(),
	}

	t.IOS.inputs.add(t, "trigger", 0)

	t.IOS.inputs.add(t, "attack", 0)
	t.IOS.inputs.add(t, "decay", 0)
	t.IOS.inputs.add(t, "release", 0)

	t.IOS.inputs.add(t, "sustain", 1)
	t.IOS.inputs.add(t, "ratioA", 0.3)
	t.IOS.inputs.add(t, "ratioDR", 0.0001)

	t.IOS.outputs.add(t, "output", NewADSROutput(t))

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
	t          *ADSR
	state      ADSRState
	lastOutput float64
}

func NewADSROutput(t *ADSR) *ADSROutput {
	return &ADSROutput{
		t:     t,
		state: Idle,
	}
}

func (o *ADSROutput) Next() float64 {
	var (
		trigger = o.t.IOS.inputs.Read("trigger")
		attack  = o.t.IOS.inputs.Read("attack") * constants.SAMPLE_RATE
		decay   = o.t.IOS.inputs.Read("decay") * constants.SAMPLE_RATE
		release = o.t.IOS.inputs.Read("release") * constants.SAMPLE_RATE
		sustain = o.t.IOS.inputs.Read("sustain")
		ratioA  = o.t.IOS.inputs.Read("ratioA")
		ratioDR = o.t.IOS.inputs.Read("ratioDR")
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

	if trigger > 0 && o.state == Idle {
		o.state = Attack
	} else if trigger <= 0 && o.state != Idle {
		o.state = Release
	}

	switch o.state {
	case Idle:
		o.lastOutput = 0
	case Attack:
		o.lastOutput = attackBase + o.lastOutput*attackCoef

		if o.lastOutput >= 1 {
			o.lastOutput = 1
			o.state = Decay
		}
	case Decay:
		o.lastOutput = decayBase + o.lastOutput*decayCoef

		if o.lastOutput <= sustain {
			o.lastOutput = sustain
			o.state = Sustain
		}
	case Sustain:
	case Release:
		o.lastOutput = releaseBase + o.lastOutput*releaseCoef

		if o.lastOutput < 0 {
			o.lastOutput = 0
			o.state = Idle
		}
	}

	return o.lastOutput
}

func calcCoef(rate float64, ratio float64) float64 {
	if rate <= 0 {
		return 0
	} else {
		return math.Exp(-math.Log((ratio+1)/ratio) / rate)
	}
}
