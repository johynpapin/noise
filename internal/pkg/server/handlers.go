package server

import (
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/internal/pkg/engine"
)

// General commands

func handleCreateTrackMessage(so socketio.Socket, player *engine.Player, m *createTrackMessage) {
	player.NewTrack(m.Name)

	emitState(so, player)
}

func handleDeleteTrackMessage(so socketio.Socket, player *engine.Player, m *deleteTrackMessage) {
	track := player.GetTrack(m.Name)

	_ = track // TODO

	emitState(so, player)
}

func handleCreateThingMessage(so socketio.Socket, player *engine.Player, m *createThingMessage) {
	track := player.GetTrack(m.TrackName)

	thing := track.ThingsManager.NewThing(m.Kind, m.Name)

	track.AddThing(thing)

	emitState(so, player)
}

func handleDeleteThingMessage(so socketio.Socket, player *engine.Player, m *deleteThingMessage) {
	track := player.GetTrack(m.TrackName)

	_ = track // TODO

	emitState(so, player)
}

func handleAttachToThingMessage(so socketio.Socket, player *engine.Player, m *attachToThingMessage) {
	track := player.GetTrack(m.TrackName)

	inputThing := track.ThingsManager.GetThing(m.InputThingName)
	outputThing := track.ThingsManager.GetThing(m.OutputThingName)

	output := outputThing.GetOutputs().Get(m.OutputName)

	inputThing.GetInputs().Attach(m.InputName, output)

	emitState(so, player)
}

func handleDetachFromThingMessage(so socketio.Socket, player *engine.Player, m *detachFromThingMessage) {
	track := player.GetTrack(m.TrackName)

	inputThing := track.ThingsManager.GetThing(m.InputThingName)

	input := inputThing.GetInputs().Inputs[m.InputName]

	input.Detach()

	emitState(so, player)
}

func handleAttachToTrackMessage(so socketio.Socket, player *engine.Player, m *attachToTrackMessage) {
	track := player.GetTrack(m.TrackName)

	thing := track.ThingsManager.GetThing(m.ThingName)

	output := thing.GetOutputs().Get(m.Name)

	track.AddOutput(output)

	emitState(so, player)
}

func handleDetachFromTrackMessage(so socketio.Socket, player *engine.Player, m *detachFromTrackMessage) {
	track := player.GetTrack(m.TrackName)

	_ = track // TODO

	emitState(so, player)
}

func handleUpdateSettingMessage(so socketio.Socket, player *engine.Player, m *updateSettingMessage) {
	track := player.GetTrack(m.TrackName)

	thing := track.ThingsManager.GetThing(m.ThingName)

	thing.GetSettings().Set(m.Name, m.Value)

	emitState(so, player)
}

// MIDI commands

func handleAttachToMIDIMessage(so socketio.Socket, player *engine.Player, m *attachToMIDIMessage) {
	track := player.GetTrack(m.TrackName)

	_ = track // TODO

	emitState(so, player)
}

func handleDetachFromMIDIMessage(so socketio.Socket, player *engine.Player, m *detachFromMIDIMessage) {
	track := player.GetTrack(m.TrackName)

	_ = track // TODO

	emitState(so, player)
}
