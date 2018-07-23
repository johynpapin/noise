package server

import (
	"encoding/json"
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/internal/pkg/engine"
	"github.com/johynpapin/noise/pkg/things"
	"sort"
)

type stateLinkedTo struct {
	ThingName  string `json:"thingName"`
	OutputName string `json:"outputName"`
}

type stateSetting struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
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
	Name     string          `json:"name"`
	X        float64         `json:"x"`
	Y        float64         `json:"y"`
	Kind     string          `json:"type"`
	Settings []*stateSetting `json:"settings"`
	Inputs   []*stateInput   `json:"inputs"`
	Outputs  []*stateOutput  `json:"outputs"`
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

func getSettings(thing things.Thing) []*stateSetting {
	var r []*stateSetting
	var keys []string

	for k := range thing.GetSettings().Settings {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		setting := thing.GetSettings().Settings[k]

		s := &stateSetting{
			Name:  k,
			Value: setting.Get(),
		}

		r = append(r, s)
	}

	return r
}

func getInputs(track *engine.Track, thing things.Thing) []*stateInput {
	var r []*stateInput
	var keys []string

	for k := range thing.GetInputs().Inputs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		input := thing.GetInputs().Inputs[k]

		i := &stateInput{
			Name: k,
		}

		if input.Output != nil {
			lt := &stateLinkedTo{
				ThingName:  track.ThingsManager.GetName(input.Output.Thing),
				OutputName: input.Output.Name,
			}

			i.LinkedTo = lt
		}

		r = append(r, i)
	}

	return r
}

func getOutputs(track *engine.Track, thing things.Thing) []*stateOutput {
	var r []*stateOutput
	var keys []string

	for k := range thing.GetOutputs().Outputs {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		o := &stateOutput{
			ThingName: track.ThingsManager.GetName(thing),
			Name:      k,
		}

		r = append(r, o)
	}

	return r
}

func getThings(track *engine.Track) []*stateThing {
	var r []*stateThing

	for _, thing := range track.Things {
		position := getPositionOfThing(thing)

		th := &stateThing{
			Name: track.ThingsManager.GetName(thing),
			X:    position.X,
			Y:    position.Y,
			Kind: track.ThingsManager.GetKind(thing),
		}

		th.Settings = getSettings(thing)
		th.Inputs = getInputs(track, thing)
		th.Outputs = getOutputs(track, thing)

		r = append(r, th)
	}

	return r
}

func generateState(e *engine.Engine) *state {
	state := &state{}

	var keys []string

	for k := range e.Player.Tracks {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		track := e.Player.Tracks[k]

		t := &stateTrack{
			Name: track.Name,
		}

		t.Things = getThings(track)

		for _, output := range track.Outputs {
			o := &stateOutput{
				ThingName: track.ThingsManager.GetName(output.Thing),
				Name:      output.Name,
			}

			t.Outputs = append(t.Outputs, o)
		}

		state.Tracks = append(state.Tracks, t)
	}

	return state
}

func emitState(so socketio.Socket, e *engine.Engine) error {
	state := generateState(e)
	stateMsg := stateMessage{*state}
	answer, _ := json.Marshal(stateMsg)

	return so.Emit("state", string(answer))
}
