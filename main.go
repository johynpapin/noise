package main

import (
	"github.com/johynpapin/noise/runtime"
	"github.com/johynpapin/noise/server"
	"github.com/johynpapin/noise/store"
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

	store := store.NewStore()

	r := runtime.NewRuntime(store)
	err := r.Start()
	if err != nil {
		log.WithError(err).Fatal("fatal error")
	}

	log.Info("list of midi devices:")
	for i, device := range r.MidiManager.GetDevices() {
		log.WithField("device", device).Info("midi device:")
		if i > 1 && device.IsInputAvailable {
			r.MidiManager.ConnectToDevice(portmidi.DeviceID(i))
		}
	}

	go server.Serve(r)

	<-stop
	r.Player.Stop()

	log.Info("Good bye!")
}
