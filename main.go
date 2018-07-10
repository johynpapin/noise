package main

import (
	"github.com/johynpapin/noise/engine"
	"github.com/johynpapin/noise/stuff"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	log.Info("== NOISE Orchestrator Is Super Enjoyable ==")

	freq := stuff.NewFloat64Generator(42)
	amp := stuff.NewFloat64Generator(1)
	offset := stuff.NewFloat64Generator(1)

	lg0 := stuff.NewLowGenSaw()

	lg0.Freq.Attach(freq)
	lg0.Amp.Attach(amp)
	lg0.Offset.Attach(offset)

	lg1 := stuff.NewLowGenTriangle()
	lg2 := stuff.NewLowGenPulse()
	lg3 := stuff.NewLowGenSaw()

	track := engine.NewTrack()

	track.AddAudible(lg0)
	track.AddAudible(lg1)
	track.AddAudible(lg2)
	track.AddAudible(lg3)

	player := engine.NewPlayer()
	player.AddTrack(track)

	stop := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			close(stop)
		}
	}()

	player.Play()
	log.Info("Playing...")

	<-stop
	player.Stop()
	log.Info("Good bye!")
}
