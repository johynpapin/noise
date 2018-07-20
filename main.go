package main

import (
	"github.com/johynpapin/noise/internal/pkg/engine"
	"github.com/johynpapin/noise/internal/pkg/server"
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

	player := engine.NewPlayer()

	go server.Serve(player)

	go player.Play()
	<-stop
	player.Stop()
	log.Info("Good bye!")
}
