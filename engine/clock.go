package engine

import (
	"fmt"
	"time"
)

type Clock struct {
	PPQN  int
	Tempo float64

	Pulses chan int
}

func NewClock(tempo float64) *Clock {
	return &Clock{
		PPQN:   24,
		Tempo:  tempo,
		Pulses: make(chan int),
	}
}

func (c *Clock) Start() {
	go func() {
		for {
			time.Sleep(c.microsecondsPerPulse())
			c.Pulses <- 1
		}
		fmt.Println("ah")
	}()
}

func (c *Clock) microsecondsPerPulse() time.Duration {
	return time.Duration((60 * 1000000000) / (float64(c.PPQN) * c.Tempo))
}
