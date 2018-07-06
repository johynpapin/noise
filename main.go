package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	log.Info("== NOISE Orchestrator Is Super Enjoyable ==")

	stereoSine := NewStereoSine(256, 540, SAMPLE_RATE)
	track := NewTrack()
	player := NewPlayer()

	track.AddAudible(stereoSine)
	player.AddTrack(track)

	patternStr := "10001110010001000100010001000100010001000100010001000100010001000"
	patternBool := make([]bool, len(patternStr))
	for i, c := range patternStr {
		if c == '1' {
			patternBool[i] = true
		}
	}

	pattern := NewPattern(patternBool)

	timer := NewTimer(120.0)
	sequencer := NewSequencer(timer, pattern)

	sequencer.AddSequencable(stereoSine)

	stop := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			close(stop)
		}
	}()

	go sequencer.Start()
	go timer.Start()
	player.Play()
	log.Info("Playing...")

	<-stop
	player.Stop()
	log.Info("Good bye!")
}
