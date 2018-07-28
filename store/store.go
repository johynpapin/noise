package store

import (
	"github.com/johynpapin/noise/engine"
	"github.com/johynpapin/noise/things"
)

type Store struct {
	idsTracks           map[int]*engine.Track
	idsInstrumentPatchs map[int]*engine.InstrumentPatch
	idsInstruments      map[int]*engine.Instrument
	idsThings           map[int]*things.Thing
	idsSettings         map[int]*things.Setting
	idsInputs           map[int]*things.Input
	idsOutputs          map[int]*things.Output
	idsLoops            map[int]*engine.Loop

	thingsIdsPositions map[int]Position

	attachedSettings []int
}

func NewStore() *Store {
	return &Store{
		idsTracks:           make(map[int]*engine.Track),
		idsInstrumentPatchs: make(map[int]*engine.InstrumentPatch),
		idsInstruments:      make(map[int]*engine.Instrument),
		idsThings:           make(map[int]*things.Thing),
		idsSettings:         make(map[int]*things.Setting),
		idsInputs:           make(map[int]*things.Input),
		idsOutputs:          make(map[int]*things.Output),
		idsLoops:            make(map[int]*engine.Loop),

		thingsIdsPositions: make(map[int]Position),
	}
}

// Get

func (s *Store) GetTrack(ID int) *engine.Track {
	return s.idsTracks[ID]
}

func (s *Store) GetInstrumentPatch(ID int) *engine.InstrumentPatch {
	return s.idsInstrumentPatchs[ID]
}

func (s *Store) GetInstrument(ID int) *engine.Instrument {
	return s.idsInstruments[ID]
}

func (s *Store) GetThing(ID int) *things.Thing {
	return s.idsThings[ID]
}

func (s *Store) GetSetting(ID int) *things.Setting {
	return s.idsSettings[ID]
}

func (s *Store) GetInput(ID int) *things.Input {
	return s.idsInputs[ID]
}

func (s *Store) GetOutput(ID int) *things.Output {
	return s.idsOutputs[ID]
}

func (s *Store) GetLoop(ID int) *engine.Loop {
	return s.idsLoops[ID]
}

// Create

func (s *Store) NewTrack(name string) *engine.Track {
	ID := len(s.idsTracks)

	track := engine.NewTrack(name)
	track.ID = ID

	s.idsTracks[ID] = track

	return track
}

func (s *Store) NewInstrumentPatch(instrumentID int) *engine.InstrumentPatch {
	ID := len(s.idsInstrumentPatchs)

	instrument := s.GetInstrument(instrumentID)

	patch := engine.NewInstrumentPatch(instrument)
	patch.ID = ID

	s.idsInstrumentPatchs[ID] = patch

	return patch
}

func (s *Store) NewInstrument(name string) *engine.Instrument {
	ID := len(s.idsInstruments)

	instrument := engine.NewInstrument(name)
	instrument.ID = ID

	s.idsInstruments[ID] = instrument

	return instrument
}

func (s *Store) NewThing(name string, kind string) *things.Thing {
	ID := len(s.idsThings)

	thing := things.NewThingFromKind(name, kind)
	thing.ID = ID

	for _, setting := range thing.Settings.Settings {
		s.storeSetting(setting)
	}

	for _, input := range thing.Inputs.Inputs {
		s.storeInput(input)
	}

	for _, output := range thing.Outputs.Outputs {
		s.storeOutput(output)
	}

	s.idsThings[ID] = thing

	return thing
}

func (s *Store) NewLoop(patchID int) *engine.Loop {
	ID := len(s.idsLoops)

	loop := engine.NewLoop(s.GetInstrumentPatch(patchID))

	loop.ID = ID

	s.idsLoops[ID] = loop

	return loop
}

func (s *Store) storeSetting(setting *things.Setting) {
	ID := len(s.idsSettings)

	setting.ID = ID

	s.idsSettings[ID] = setting
}

func (s *Store) storeInput(input *things.Input) {
	ID := len(s.idsInputs)

	input.ID = ID

	s.idsInputs[ID] = input
}

func (s *Store) storeOutput(output *things.Output) {
	ID := len(s.idsOutputs)

	output.ID = ID

	s.idsOutputs[ID] = output
}

// Attach

func (s *Store) AddInstrumentPatchToTrack(patchID int, trackID int) {
	s.GetTrack(trackID).AddInstrumentPatch(s.GetInstrumentPatch(patchID))
}

func (s *Store) AddThingToInstrument(thingID int, instrumentID int) {
	s.GetInstrument(instrumentID).AddThing(s.GetThing(thingID))
}

func (s *Store) AttachOutputToInput(inputID int, outputID int) {
	s.GetInput(inputID).Attach(s.GetOutput(outputID))
}

func (s *Store) AttachOutputToInstrument(outputID int, instrumentID int) {
	s.GetInstrument(instrumentID).SetExitPoint(s.GetOutput(outputID))
}

func (s *Store) AttachInputToInstrument(inputID int, instrumentID int) {
	s.GetInput(inputID).Attach(s.GetInstrument(instrumentID).EntryPoint)
}

func (s *Store) AttachSettingToMIDI(midiManager *engine.MidiManager, patchID int, settingID int) {
	midiManager.AttachSettingToNextEvent(s.GetInstrumentPatch(patchID), s.GetSetting(settingID))
	s.attachedSettings = append(s.attachedSettings, settingID)
}

// Update

func (s *Store) UpdateThingPosition(thingID int, x float64, y float64) {
	s.thingsIdsPositions[thingID] = Position{X: x, Y: y}
}

func (s *Store) UpdateSetting(settingID int, name string, value float64, min float64, max float64) {
	setting := s.GetSetting(settingID)
	setting.ExposedName = name
	setting.Value = value
	setting.Min = min
	setting.Max = max
}

func (s *Store) UpdateExposedSetting(patchID int, settingID int, value float64) {
	s.GetInstrumentPatch(patchID).Set(s.GetSetting(settingID), value)
}

func (s *Store) ExposeSetting(instrumentID int, settingID int, name string) {
	s.GetInstrument(instrumentID).ExposeSetting(name, s.GetSetting(settingID))
}

func (s *Store) HideSetting(instrumentID int, settingID int) {
	s.GetInstrument(instrumentID).HideSetting(s.GetSetting(settingID))
}

func (s *Store) StartRecording(loopID int) {
	loop := s.GetLoop(loopID)

	loop.StartRecording()
}

func (s *Store) StopRecording(loopID int) {
	loop := s.GetLoop(loopID)

	loop.StopRecording()
}
