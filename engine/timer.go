package engine

import (
	"github.com/johynpapin/noise/constants"
	"time"
)

type Timer struct {
	Pulses chan int
	tempo  float32
}

func NewTimer(tempo float32) *Timer {
	return &Timer{make(chan int), tempo}
}

func (t *Timer) Start() {
	for {
		time.Sleep(t.MicrosecondsPerPulse())
		t.Pulses <- 1
	}
}

func (t *Timer) MicrosecondsPerPulse() time.Duration {
	return time.Duration((60.0 * 1000000000.0) / (constants.PPQN * t.tempo))
}
