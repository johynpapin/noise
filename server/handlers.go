package server

import (
	"github.com/googollee/go-socket.io"
	"github.com/johynpapin/noise/runtime"
)

type commandHandlerFunc func(socketio.Socket, *runtime.Runtime, *commandMessage)

var commandHandlers = map[string]commandHandlerFunc{
	"createTrack": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		track := e.Store.NewTrack(m.Name)

		e.Player.AddTrack(track)
	},
	"updateTrack": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"deleteTrack": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"createInstrument": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.NewInstrument(m.Name)
	},
	"updateInstrument": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"deleteInstrument": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"addInstrumentToTrack": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		patch := e.Store.NewInstrumentPatch(m.InstrumentID)

		e.Store.AddInstrumentPatchToTrack(patch.ID, m.TrackID)
	},
	"createThing": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		thing := e.Store.NewThing(m.Name, m.Kind)

		e.Store.AddThingToInstrument(thing.ID, m.InstrumentID)
	},
	"updateThing": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"deleteThing": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"attachToThing": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.AttachOutputToInput(m.InputID, m.OutputID)
	},
	"detachFromThing": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"attachToInstrumentEntryPoint": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.AttachInputToInstrument(m.InputID, m.InstrumentID)
	},
	"detachFromInstrumentEntryPoint": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"attachToInstrumentExitPoint": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.AttachOutputToInstrument(m.OutputID, m.InstrumentID)
	},
	"detachFromInstrumentExitPoint": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"updateSetting": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.UpdateSetting(m.SettingID, m.Name, m.Value, m.Min, m.Max)
	},
	"updateExposedSetting": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.UpdateExposedSetting(m.InstrumentPatchID, m.SettingID, m.Value)
	},
	"attachToMIDI": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.AttachSettingToMIDI(e.MidiManager, m.InstrumentPatchID, m.SettingID)
	},
	"detachFromMIDI": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
	},
	"updateThingPosition": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.UpdateThingPosition(m.ThingID, m.X, m.Y)
	},
	"updateCurrentInstrumentPatch": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.MidiManager.SetCurrentInstrumentPatch(e.Store.GetInstrumentPatch(m.InstrumentPatchID))
	},
	"exposeSetting": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.ExposeSetting(m.InstrumentID, m.SettingID, m.Name)
	},
	"hideSetting": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.HideSetting(m.InstrumentID, m.SettingID)
	},
	"createLoop": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		loop := e.Store.NewLoop(m.InstrumentPatchID)

		e.Sequencer.AddLoop(loop)
	},
	"startRecording": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.StartRecording(m.LoopID)
		e.MidiManager.SetCurrentLoop(e.Store.GetLoop(m.LoopID))
	},
	"stopRecording": func(so socketio.Socket, e *runtime.Runtime, m *commandMessage) {
		e.Store.StopRecording(m.LoopID)
	},
}
