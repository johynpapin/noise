package engine

import (
	"github.com/rakyll/portmidi"
)

type MidiDevice struct {
	deviceID portmidi.DeviceID
	stream   *portmidi.Stream
}

func NewMidiDevice(deviceID portmidi.DeviceID) *MidiDevice {
	return &MidiDevice{
		deviceID: deviceID,
	}
}

func (md *MidiDevice) Start() error {
	var err error
	md.stream, err = portmidi.NewInputStream(md.deviceID, 1024)
	if err != nil {
		return err
	}

	ch := md.stream.Listen()
	for event := range ch {
		_ = event
	}

	return nil
}

func (md *MidiDevice) Stop() error {
	return md.stream.Close()
}
