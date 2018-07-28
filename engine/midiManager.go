package engine

import (
	"github.com/gomidi/midi"
	"github.com/gomidi/midi/midimessage/channel"
	"github.com/johynpapin/noise/things"
	"github.com/rakyll/portmidi"
	"math"
)

type attachedSetting struct {
	InstrumentPatch *InstrumentPatch
	Setting         *things.Setting
}

type MidiManager struct {
	devices          []*MidiDevice
	controlChangeMap map[[2]uint8]*attachedSetting

	currentInstrumentPatch *InstrumentPatch
	waitingSetting         *attachedSetting

	sequencer   *Sequencer
	currentLoop *Loop
}

func NewMidiManager() *MidiManager {
	return &MidiManager{
		controlChangeMap: make(map[[2]uint8]*attachedSetting),
	}
}

func (mm *MidiManager) Start() error {
	err := portmidi.Initialize()
	if err != nil {
		return err
	}

	return nil
}

func (mm *MidiManager) Stop() error {
	err := portmidi.Terminate()
	if err != nil {
		return err
	}

	return nil
}

func (mm *MidiManager) Process(msg midi.Message) {
	switch m := msg.(type) {
	case channel.NoteOn:
		if mm.currentInstrumentPatch != nil {
			if mm.currentLoop != nil {
				mm.currentLoop.SaveMessage(msg)
			}

			if m.Velocity() == 0 {
				mm.currentInstrumentPatch.NoteOff(midiToHz(float64(m.Key())))
			} else {
				mm.currentInstrumentPatch.NoteOn(midiToHz(float64(m.Key())))
			}
		}
	case channel.NoteOff:
		if mm.currentLoop != nil {
			mm.currentLoop.SaveMessage(msg)
		}

		if mm.currentInstrumentPatch != nil {
			mm.currentInstrumentPatch.NoteOff(midiToHz(float64(m.Key())))
		}
	case channel.ControlChange:
		if mm.waitingSetting != nil {
			mm.controlChangeMap[[2]uint8{m.Channel(), m.Controller()}] = mm.waitingSetting
			mm.waitingSetting = nil
		}

		setting, ok := mm.controlChangeMap[[2]uint8{m.Channel(), m.Controller()}]
		if ok {
			slope := (setting.Setting.Max - setting.Setting.Min) / 127
			setting.InstrumentPatch.Set(setting.Setting, setting.Setting.Min+slope*float64(m.Value()))
		}
	}
}

func (mm *MidiManager) ProcessLoopMessage(patch *InstrumentPatch, msg midi.Message) {
	switch m := msg.(type) {
	case channel.NoteOn:
		patch.NoteOn(midiToHz(float64(m.Key())))
	case channel.NoteOff:
		patch.NoteOff(midiToHz(float64(m.Key())))
	}
}

func (mm *MidiManager) ConnectToDevice(deviceID portmidi.DeviceID) {
	md := NewMidiDevice(deviceID)

	mm.devices = append(mm.devices, md)

	go func() {
		for msg := range md.Messages {
			mm.Process(msg)
		}
	}()

	md.Start()
}

func (mm *MidiManager) AttachSettingToNextEvent(patch *InstrumentPatch, setting *things.Setting) {
	mm.waitingSetting = &attachedSetting{
		InstrumentPatch: patch,
		Setting:         setting,
	}
}

func (mm *MidiManager) GetDevices() []*portmidi.DeviceInfo {
	var devices []*portmidi.DeviceInfo

	nbDevices := portmidi.CountDevices()

	for i := 0; i < nbDevices; i++ {
		devices = append(devices, portmidi.Info(portmidi.DeviceID(i)))
	}

	return devices
}

func (mm *MidiManager) SetCurrentInstrumentPatch(patch *InstrumentPatch) {
	mm.currentInstrumentPatch = patch
}

func (mm *MidiManager) SetCurrentLoop(loop *Loop) {
	mm.currentLoop = loop
}

func midiToHz(key float64) float64 {
	return 440 * math.Pow(2, (key-69)/12)
}
