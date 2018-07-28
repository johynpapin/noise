package store

import (
	"sort"
)

type StateSetting struct {
	ID          int     `json:"ID"`
	Name        string  `json:"name"`
	ExposedName string  `json:"exposedName"`
	Min         float64 `json:"min"`
	Max         float64 `json:"max"`
	Value       float64 `json:"value"`
	Attached    bool    `json:"attached"`
}

type StateLinkedTo struct {
	OutputID int `json:"outputID"`
	ThingID  int `json:"thingID"`
}

type StateInput struct {
	ID       int            `json:"ID"`
	Name     string         `json:"name"`
	LinkedTo *StateLinkedTo `json:"linkedTo"`
}

type StateOutput struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type StateThing struct {
	ID       int             `json:"ID"`
	Name     string          `json:"name"`
	Settings []*StateSetting `json:"settings"`
	Inputs   []*StateInput   `json:"inputs"`
	Outputs  []*StateOutput  `json:"outputs"`
	X        float64         `json:"x"`
	Y        float64         `json:"y"`
}

type StateLoop struct {
	ID                int  `json:"ID"`
	IsRecording       bool `json:"isRecording"`
	InstrumentPatchID int  `json:"instrumentPatchID"`
}

type StateInstrumentPatch struct {
	ID int `json:"ID"`

	Instrument *StateInstrument `json:"instrument"`
	Settings   []*StateSetting  `json:"settings"`
}

type StateInstrument struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`

	Things []*StateThing `json:"things"`

	EntryPoint *StateOutput `json:"entryPoint"`
	ExitPoint  *StateInput  `json:"exitPoint"`

	Settings []*StateSetting `json:"settings"`
}

type StateTrack struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`

	InstrumentsPatchs []*StateInstrumentPatch `json:"instrumentPatchs"`
}

type State struct {
	Tracks      []*StateTrack      `json:"tracks"`
	Instruments []*StateInstrument `json:"instruments"`
	Loops       []*StateLoop       `json:"loops"`
}

func (s *Store) exportSettingsOfThing(thingID int) []*StateSetting {
	thing := s.GetThing(thingID)

	var r []*StateSetting

	var keys []string

	for key := range thing.Settings.Settings {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, settingName := range keys {
		setting := thing.Settings.Settings[settingName]

		r = append(r, &StateSetting{
			ID:          setting.ID,
			Name:        settingName,
			ExposedName: setting.ExposedName,
			Min:         setting.Min,
			Max:         setting.Max,
			Value:       setting.Value,
		})
	}

	return r
}

func (s *Store) exportInputsOfThing(thingID int) []*StateInput {
	thing := s.GetThing(thingID)

	var r []*StateInput

	var keys []string

	for key := range thing.Inputs.Inputs {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, inputName := range keys {
		input := thing.Inputs.Inputs[inputName]
		inputID := input.ID

		stateInput := &StateInput{
			ID:   inputID,
			Name: inputName,
		}

		if input.Output != nil {
			if input.Output.Name == "Entry Point" {
				stateInput.LinkedTo = &StateLinkedTo{
					OutputID: -20,
					ThingID:  -2,
				}
			} else {
				stateInput.LinkedTo = &StateLinkedTo{
					OutputID: input.Output.ID,
					ThingID:  input.Output.Thing.ID,
				}
			}
		}

		r = append(r, stateInput)
	}

	return r
}

func (s *Store) exportOutputsOfThing(thingID int) []*StateOutput {
	thing := s.GetThing(thingID)

	var r []*StateOutput

	var keys []string

	for key := range thing.Outputs.Outputs {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, outputName := range keys {
		output := thing.Outputs.Outputs[outputName]
		outputID := output.ID

		r = append(r, &StateOutput{
			ID:   outputID,
			Name: outputName,
		})
	}

	return r
}

func (s *Store) exportThingsOfInstrument(instrumentID int) []*StateThing {
	instrument := s.GetInstrument(instrumentID)

	var r []*StateThing

	for _, thing := range instrument.Things {
		r = append(r, &StateThing{
			ID:   thing.ID,
			Name: thing.Name,

			Settings: s.exportSettingsOfThing(thing.ID),
			Inputs:   s.exportInputsOfThing(thing.ID),
			Outputs:  s.exportOutputsOfThing(thing.ID),

			X: s.thingsIdsPositions[thing.ID].X,
			Y: s.thingsIdsPositions[thing.ID].Y,
		})
	}

	return r
}

func (s *Store) exportSettingsOfInstrumentPatch(patchID int) []*StateSetting {
	var r []*StateSetting

	patch := s.GetInstrumentPatch(patchID)

	for _, setting := range patch.Instrument.Settings {
		attached := false

		for _, setID := range s.attachedSettings {
			if setID == setting.ID {
				attached = true
				break
			}
		}

		r = append(r, &StateSetting{
			ID:          setting.ID,
			Name:        setting.Name,
			ExposedName: setting.ExposedName,
			Value:       patch.Get(setting),
			Min:         setting.Min,
			Max:         setting.Max,
			Attached:    attached,
		})
	}

	return r
}

func (s *Store) exportInstrumentPatchsOfTrack(trackID int) []*StateInstrumentPatch {
	var r []*StateInstrumentPatch

	track := s.GetTrack(trackID)

	for _, patch := range track.InstrumentPatchs {
		r = append(r, &StateInstrumentPatch{
			ID: patch.ID,
			Instrument: &StateInstrument{
				Name:     patch.Instrument.Name,
				Settings: s.exportSettingsOfInstrument(patch.Instrument.ID),
			},
			Settings: s.exportSettingsOfInstrumentPatch(patch.ID),
		})
	}

	return r
}

func (s *Store) exportTracks() []*StateTrack {
	var r []*StateTrack

	var keys []int

	for key := range s.idsTracks {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, trackID := range keys {
		track := s.GetTrack(trackID)

		r = append(r, &StateTrack{
			ID:   trackID,
			Name: track.Name,

			InstrumentsPatchs: s.exportInstrumentPatchsOfTrack(trackID),
		})
	}

	return r
}

func (s *Store) exportLoops() []*StateLoop {
	var r []*StateLoop

	var keys []int

	for key := range s.idsLoops {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, loopID := range keys {
		loop := s.GetLoop(loopID)

		r = append(r, &StateLoop{
			ID:                loopID,
			IsRecording:       loop.IsRecording,
			InstrumentPatchID: loop.InstrumentPatch.ID,
		})
	}

	return r
}

func (s *Store) exportSettingsOfInstrument(instrumentID int) []*StateSetting {
	instrument := s.GetInstrument(instrumentID)

	var r []*StateSetting

	for _, setting := range instrument.Settings {
		r = append(r, &StateSetting{
			ID:          setting.ID,
			Name:        setting.Name,
			ExposedName: setting.ExposedName,
			Min:         setting.Min,
			Max:         setting.Max,
			Value:       setting.Value,
		})
	}

	return r
}

func (s *Store) exportInstruments() []*StateInstrument {
	var r []*StateInstrument

	var keys []int

	for key := range s.idsInstruments {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	for _, instrumentID := range keys {
		instrument := s.GetInstrument(instrumentID)

		stateInstrument := &StateInstrument{
			ID:   instrumentID,
			Name: instrument.Name,

			Things: s.exportThingsOfInstrument(instrumentID),

			EntryPoint: &StateOutput{
				ID:   -20,
				Name: "Entry Point",
			},
			ExitPoint: &StateInput{
				ID:   -30,
				Name: "Exit Point",
			},
		}

		if instrument.ExitPoint != nil {
			outputID := instrument.ExitPoint.ID

			stateInstrument.ExitPoint.LinkedTo = &StateLinkedTo{
				OutputID: outputID,
				ThingID:  instrument.ExitPoint.Thing.ID,
			}
		}

		r = append(r, stateInstrument)
	}

	return r
}

func (s *Store) ExportState() *State {
	return &State{
		Tracks:      s.exportTracks(),
		Instruments: s.exportInstruments(),
		Loops:       s.exportLoops(),
	}
}
