package server

import (
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/runtime"
	"github.com/johynpapin/noise/things"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Serve launches an http server and a socket.io server. This makes the NOISE web panel work.
func Serve(r *runtime.Runtime) {
	mux := http.NewServeMux()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		log.Info("A panel is now connected.")

		emitState(so, r)

		thingsMsg := thingsMessage{things.Kinds}
		answer, _ := json.Marshal(thingsMsg)
		so.Emit("things", string(answer))

		ticker := time.NewTicker(10 * time.Millisecond)

		go func() {
			for _ = range ticker.C {
				err := emitState(so, r)
				if err != nil {
					ticker.Stop()
					return
				}
			}
		}()

		// General commands

		so.On("command", func(msg string) {
			m := &commandMessage{}
			err := json.Unmarshal([]byte(msg), m)
			if err != nil {
				log.WithField("error", err).Error("error reading message")
				return
			}

			log.WithField("command", m.Command).Info("command received")

			commandHandlers[m.Command](so, r, m)
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	mux.Handle("/socket.io/", server)
	mux.Handle("/", http.FileServer(http.Dir("./assets/panel/dist")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Info("Serving at localhost:4242...")
	log.Fatal(http.ListenAndServe(":4242", handler))
}
