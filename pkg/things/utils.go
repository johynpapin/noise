package things

import (
	"github.com/johynpapin/noise/pkg/constants"
	"math"
)

func StepPhase(phase, freq float64) float64 {
	phase += freq * 2 * math.Pi / constants.SAMPLE_RATE

	for phase >= 2*math.Pi {
		phase -= math.Pi
	}

	return phase
}
