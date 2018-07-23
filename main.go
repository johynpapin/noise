package main

import (
	"github.com/johynpapin/noise/internal/pkg/engine"
	"github.com/johynpapin/noise/internal/pkg/server"
	"github.com/rakyll/portmidi"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	log.Info("== NOISE Orchestrator Is Super Enjoyable ==")

	stop := make(chan bool)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			close(stop)
		}
	}()

	e := engine.NewEngine()
	err := e.Start()
	if err != nil {
		log.WithError(err).Fatal("fatal error")
	}

	log.Info("list of midi devices:")
	for _, device := range e.MidiManager.GetDevices() {
		log.WithField("device", device).Info("midi device:")
	}

	e.MidiManager.ConnectToDevice(portmidi.DeviceID(3))

	go server.Serve(e)

	<-stop
	e.Player.Stop()

	log.Info("Good bye!")
}
