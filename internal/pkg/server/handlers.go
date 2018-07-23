package server

import (
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/internal/pkg/engine"
)

// General commands

func handleCreateTrackMessage(so socketio.Socket, e *engine.Engine, m *createTrackMessage) {
	e.Player.NewTrack(m.Name)
}

func handleDeleteTrackMessage(so socketio.Socket, e *engine.Engine, m *deleteTrackMessage) {
	track := e.Player.GetTrack(m.Name)

	_ = track // TODO
}

func handleCreateThingMessage(so socketio.Socket, e *engine.Engine, m *createThingMessage) {
	track := e.Player.GetTrack(m.TrackName)

	thing := track.ThingsManager.NewThing(m.Kind, m.Name)

	track.AddThing(thing)
}

func handleDeleteThingMessage(so socketio.Socket, e *engine.Engine, m *deleteThingMessage) {
	track := e.Player.GetTrack(m.TrackName)

	_ = track // TODO
}

func handleAttachToThingMessage(so socketio.Socket, e *engine.Engine, m *attachToThingMessage) {
	track := e.Player.GetTrack(m.TrackName)

	inputThing := track.ThingsManager.GetThing(m.InputThingName)
	outputThing := track.ThingsManager.GetThing(m.OutputThingName)

	output := outputThing.GetOutputs().Get(m.OutputName)

	inputThing.GetInputs().Attach(m.InputName, output)
}

func handleDetachFromThingMessage(so socketio.Socket, e *engine.Engine, m *detachFromThingMessage) {
	track := e.Player.GetTrack(m.TrackName)

	inputThing := track.ThingsManager.GetThing(m.InputThingName)

	input := inputThing.GetInputs().Inputs[m.InputName]

	input.Detach()
}

func handleAttachToTrackMessage(so socketio.Socket, e *engine.Engine, m *attachToTrackMessage) {
	track := e.Player.GetTrack(m.TrackName)

	thing := track.ThingsManager.GetThing(m.ThingName)

	output := thing.GetOutputs().Get(m.Name)

	track.AddOutput(output)
}

func handleDetachFromTrackMessage(so socketio.Socket, e *engine.Engine, m *detachFromTrackMessage) {
	track := e.Player.GetTrack(m.TrackName)

	_ = track // TODO
}

func handleUpdateSettingMessage(so socketio.Socket, e *engine.Engine, m *updateSettingMessage) {
	track := e.Player.GetTrack(m.TrackName)

	thing := track.ThingsManager.GetThing(m.ThingName)

	thing.GetSettings().Set(m.Name, m.Value)
}

// MIDI commands

func handleAttachToMIDIMessage(so socketio.Socket, e *engine.Engine, m *attachToMIDIMessage) {
	track := e.Player.GetTrack(m.TrackName)

	thing := track.ThingsManager.GetThing(m.ThingName)

	setting := thing.GetSettings().Settings[m.Name]

	e.MidiManager.AttachSettingToNextEvent(setting)
}

func handleDetachFromMIDIMessage(so socketio.Socket, e *engine.Engine, m *detachFromMIDIMessage) {
	track := e.Player.GetTrack(m.TrackName)

	_ = track // TODO
}

// UI commands

func handleUpdateThingsPositionMessage(so socketio.Socket, e *engine.Engine, m *updateThingsPositionMessage) {
	track := e.Player.GetTrack(m.TrackName)

	thing := track.ThingsManager.GetThing(m.Name)

	storePositionOfThing(thing, position{X: m.X, Y: m.Y})
}
