package stuff

import (
	"github.com/johynpapin/noise/constants"
	"math"
)

type LowGenSaw struct {
	Freq, Amp, Offset *Input
	phase             float64
}

func NewLowGenSaw() *LowGenSaw {
	return &LowGenSaw{
		Freq:   NewInput(),
		Amp:    NewInput(),
		Offset: NewInput(),
	}
}

func (s *LowGenSaw) Next() float64 {
	var (
		freq   = s.Freq.Next()
		amp    = s.Amp.Next()
		offset = s.Offset.Next()
	)

	next := (2*s.phase)/(2*math.Pi) - 1

	_, s.phase = math.Modf(s.phase + freq/constants.SAMPLE_RATE)

	return amp*next + offset
}
