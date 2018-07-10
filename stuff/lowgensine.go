package stuff

import (
	"github.com/johynpapin/noise/constants"
	"math"
)

type LowGenSine struct {
	Freq, Amp, Offset *Input
	phase             float64
}

func NewLowGenSine() *LowGenSine {
	return &LowGenSine{
		Freq:   NewInput(),
		Amp:    NewInput(),
		Offset: NewInput(),
	}
}

func (s *LowGenSine) Next() float64 {
	var (
		freq   = s.Freq.Next()
		amp    = s.Amp.Next()
		offset = s.Offset.Next()
	)

	next := math.Sin(2 * math.Pi * s.phase)

	_, s.phase = math.Modf(s.phase + freq/constants.SAMPLE_RATE)

	return amp*next + offset
}
