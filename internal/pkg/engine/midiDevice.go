package engine

import (
	"bytes"
	"github.com/gomidi/midi"
	"github.com/gomidi/midi/midimessage/realtime"
	"github.com/gomidi/midi/midireader"
	"github.com/rakyll/portmidi"
	log "github.com/sirupsen/logrus"
	"io"
)

type MidiDevice struct {
	deviceID portmidi.DeviceID
	stream   *portmidi.Stream
	buffer   *bytes.Buffer
	reader   midi.Reader
	Messages chan midi.Message
}

func NewMidiDevice(deviceID portmidi.DeviceID) *MidiDevice {
	md := &MidiDevice{
		deviceID: deviceID,
		buffer:   &bytes.Buffer{},
		Messages: make(chan midi.Message),
	}

	md.reader = midireader.New(md.buffer, md.writeRealtime)

	return md
}

func (md *MidiDevice) writeRealtime(m realtime.Message) {
	log.Println("%s\n", m.String())
}

func (md *MidiDevice) Start() error {
	var err error
	md.stream, err = portmidi.NewInputStream(md.deviceID, 1024)
	if err != nil {
		return err
	}

	ch := md.stream.Listen()

	go func() {
		for event := range ch {
			md.buffer.Write([]byte{byte(event.Status), byte(event.Data1), byte(event.Data2)})

			var errRead error
			var msg midi.Message

			for {
				msg, errRead = md.reader.Read()
				if errRead != nil {
					break
				}

				md.Messages <- msg
			}

			if errRead != nil && errRead != io.EOF {
				log.WithField("error", err).Error("Error: %s\n", errRead.Error())
			}

			md.buffer.Reset()
		}
	}()

	return nil
}

func (md *MidiDevice) Stop() error {
	return md.stream.Close()
}
