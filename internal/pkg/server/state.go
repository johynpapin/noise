package server

import (
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/internal/pkg/engine"
	log "github.com/sirupsen/logrus"
)

type stateLinkedTo struct {
	ThingName  string `json:"thingName"`
	OutputName string `json:"outputName"`
}

type stateInput struct {
	Name     string         `json:"name"`
	LinkedTo *stateLinkedTo `json:"linkedTo"`
}

type stateOutput struct {
	ThingName string `json:"thingName"`
	Name      string `json:"name"`
}

type stateThing struct {
	Name    string         `json:"name"`
	X       float64        `json:"x"`
	Y       float64        `json:"y"`
	Kind    string         `json:"type"`
	Inputs  []*stateInput  `json:"inputs"`
	Outputs []*stateOutput `json:"outputs"`
}

type stateAudible struct {
	ThingName  string `json:"thingName"`
	OutputName string `json:"outputName"`
}

type stateTrack struct {
	Name    string         `json:"name"`
	Things  []*stateThing  `json:"things"`
	Outputs []*stateOutput `json:"outputs"`
}

type state struct {
	Tracks []*stateTrack `json:"tracks"`
}

func generateState(player *engine.Player) *state {
	state := &state{}

	for _, track := range player.Tracks {
		t := stateTrack{
			Name: track.Name,
		}

		log.WithField("track", t).Info("new track")

		for _, thing := range track.Things {
			s := stateThing{
				Name: track.ThingsManager.GetName(thing),
				X:    0,
				Y:    0,
				Kind: track.ThingsManager.GetKind(thing),
			}

			log.WithField("thing", s).Info("\tnew thing")

			for inputName, input := range thing.GetInputs().Inputs {
				i := stateInput{
					Name: inputName,
				}

				log.WithField("input", i).Info("\t\tnew input")

				if input.Output != nil {
					lt := stateLinkedTo{
						ThingName:  track.ThingsManager.GetName(input.Output.Thing),
						OutputName: input.Output.Name,
					}

					log.WithField("linkedTo", lt).Info("\t\t\tnew linkedTo")

					i.LinkedTo = &lt
				}

				s.Inputs = append(s.Inputs, &i)
			}

			for outputName := range thing.GetOutputs().Outputs {
				o := stateOutput{
					ThingName: track.ThingsManager.GetName(thing),
					Name:      outputName,
				}

				log.WithField("output", o).Info("\t\tnew output")

				s.Outputs = append(s.Outputs, &o)
			}

			t.Things = append(t.Things, &s)
		}

		for _, output := range track.Outputs {
			o := stateOutput{
				ThingName: track.ThingsManager.GetName(output.Thing),
				Name:      output.Name,
			}

			log.WithField("output", o).Info("\tnew output")

			t.Outputs = append(t.Outputs, &o)
		}

		state.Tracks = append(state.Tracks, &t)
	}

	return state
}

func emitState(so socketio.Socket, player *engine.Player) {
	state := generateState(player)
	stateMsg := stateMessage{*state}
	answer, _ := json.Marshal(stateMsg)

	so.Emit("state", string(answer))
}
