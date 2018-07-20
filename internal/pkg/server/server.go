package server

import (
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/internal/pkg/engine"
	"github.com/johynpapin/noise/pkg/things"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Serve launches an http server and a socket.io server. This makes the NOISE web panel work.
func Serve(player *engine.Player) {
	mux := http.NewServeMux()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Info("A panel is now connected.")

		state := generateState(player)
		stateMsg := stateMessage{*state}
		answer, _ := json.Marshal(stateMsg)
		so.Emit("state", string(answer))

		thingsMsg := thingsMessage{things.Kinds}
		answer, _ = json.Marshal(thingsMsg)
		so.Emit("things", string(answer))

		// General commands

		so.On("createTrack", func(msg string) {
			m := &createTrackMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleCreateTrackMessage(so, player, m)
		})

		so.On("deleteTrack", func(msg string) {
			m := &deleteTrackMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleDeleteTrackMessage(so, player, m)
		})

		so.On("createThing", func(msg string) {
			m := &createThingMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleCreateThingMessage(so, player, m)
		})

		so.On("deleteThing", func(msg string) {
			m := &deleteThingMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleDeleteThingMessage(so, player, m)
		})

		so.On("attachToThing", func(msg string) {
			m := &attachToThingMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleAttachToThingMessage(so, player, m)
		})

		so.On("detachFromThing", func(msg string) {
			m := &detachFromThingMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleDetachFromThingMessage(so, player, m)
		})

		so.On("attachToTrack", func(msg string) {
			m := &attachToTrackMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleAttachToTrackMessage(so, player, m)
		})

		so.On("detachFromTrack", func(msg string) {
			m := &detachFromTrackMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleDetachFromTrackMessage(so, player, m)
		})

		so.On("updateSettingMessage", func(msg string) {
			m := &updateSettingMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleUpdateSettingMessage(so, player, m)
		})

		// MIDI commands

		so.On("attachToMIDI", func(msg string) {
			m := &attachToMIDIMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleAttachToMIDIMessage(so, player, m)
		})

		so.On("detachFromMIDI", func(msg string) {
			m := &detachFromMIDIMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			handleDetachFromMIDIMessage(so, player, m)
		})

		so.On("disconnection", func() {
			log.Println("A panel is now disconnected.")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	mux.Handle("/socket.io/", server)
	mux.Handle("/", http.FileServer(http.Dir("./assets")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Info("Serving at localhost:4242...")
	log.Fatal(http.ListenAndServe(":4242", handler))
}