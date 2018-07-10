package stuff

import (
	"github.com/johynpapin/noise/constants"
	"math"
)

type LowGenTriangle struct {
	Freq, Amp, Offset *Input
	phase             float64
}

func NewLowGenTriangle() *LowGenTriangle {
	return &LowGenTriangle{
		Freq:   NewInput(),
		Amp:    NewInput(),
		Offset: NewInput(),
	}
}

func (s *LowGenTriangle) Next() float64 {
	var (
		freq   = s.Freq.Next()
		amp    = s.Amp.Next()
		offset = s.Offset.Next()
		next   float64
	)

	if s.phase < math.Pi {
		next = 2*math.Pi*s.phase - 1
	} else {
		next = 3 - 2*math.Pi*s.phase
	}

	_, s.phase = math.Modf(s.phase + freq/constants.SAMPLE_RATE)

	return amp*next + offset
}
