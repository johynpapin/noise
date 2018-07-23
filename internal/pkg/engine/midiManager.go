package engine

import (
	"github.com/gomidi/midi"
	"github.com/gomidi/midi/midimessage/channel"
	"github.com/johynpapin/noise/pkg/things"
	"github.com/rakyll/portmidi"
	log "github.com/sirupsen/logrus"
)

type MidiManager struct {
	devices          []*MidiDevice
	controlChangeMap map[[2]uint8]*things.Setting

	waitingSetting *things.Setting
}

func NewMidiManager() *MidiManager {
	return &MidiManager{
		controlChangeMap: make(map[[2]uint8]*things.Setting),
	}
}

func (mm *MidiManager) Start() error {
	err := portmidi.Initialize()
	if err != nil {
		return err
	}

	go func() {

	}()

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
		log.WithFields(log.Fields{
			"channel":  m.Channel(),
			"key":      m.Key(),
			"velocity": m.Velocity(),
		}).Info("note on")
	case channel.NoteOff:
		log.WithFields(log.Fields{
			"channel": m.Channel(),
			"key":     m.Key(),
		}).Info("note off")
	case channel.ControlChange:
		if mm.waitingSetting != nil {
			mm.controlChangeMap[[2]uint8{m.Channel(), m.Controller()}] = mm.waitingSetting
			mm.waitingSetting = nil
		}

		setting, ok := mm.controlChangeMap[[2]uint8{m.Channel(), m.Controller()}]
		if ok {
			setting.Set(float64(m.Value()))
		}

		log.WithFields(log.Fields{
			"channel":    m.Channel(),
			"controller": m.Controller(),
			"value":      m.Value(),
		}).Info("control change")
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

func (mm *MidiManager) AttachSettingToNextEvent(setting *things.Setting) {
	mm.waitingSetting = setting
}

func (mm *MidiManager) GetDevices() []*portmidi.DeviceInfo {
	var devices []*portmidi.DeviceInfo

	nbDevices := portmidi.CountDevices()

	for i := 0; i < nbDevices; i++ {
		devices = append(devices, portmidi.Info(portmidi.DeviceID(i)))
	}

	return devices
}
