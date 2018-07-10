package stuff

import (
	"github.com/johynpapin/noise/constants"
	"math"
)

type LowGenPulse struct {
	Freq, Amp, Pw, Offset *Input
	phase                 float64
}

func NewLowGenPulse() *LowGenPulse {
	return &LowGenPulse{
		Freq:   NewInput(),
		Amp:    NewInput(),
		Pw:     NewInput(),
		Offset: NewInput(),
	}
}

func (s *LowGenPulse) Next() float64 {
	var (
		freq   = s.Freq.Next()
		amp    = s.Amp.Next()
		pw     = s.Pw.Next()
		offset = s.Offset.Next()
		next   float64
	)

	if s.phase < math.Pi*pw {
		next = 1
	} else {
		next = -1
	}

	_, s.phase = math.Modf(s.phase + freq/constants.SAMPLE_RATE)

	return amp*next + offset
}
