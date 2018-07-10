package main

import (
	"github.com/rakyll/portmidi"
	log "github.com/sirupsen/logrus"
	"github.com/xlab/midievent"
	"math"
)

type MidiMapper struct {
	stream    *portmidi.Stream
	playables []Playable
}

func NewMidiMapper() *MidiMapper {
	return &MidiMapper{nil, nil}
}

func (mm *MidiMapper) Start() error {
	err := portmidi.Initialize()
	if err != nil {
		return err
	}

	log.Info("midi devices:")
	for i := 0; i < portmidi.CountDevices(); i++ {
		log.WithFields(log.Fields{
			"id":   i,
			"info": portmidi.Info(portmidi.DeviceID(i)),
		}).Info("device:")
	}

	mm.stream, err = portmidi.NewInputStream(portmidi.DeviceID(3), 1024)
	if err != nil {
		return err
	}

	ch := mm.stream.Listen()
	for event := range ch {
		log.WithField("event", event).Info("midi event:")
		if midievent.IsNoteOn(midievent.Event(event.Status)) {
			n := 440 * math.Pow(2, (float64(event.Data1)-69)/12)
			for _, p := range mm.playables {
				if event.Data2 == 0 {
					p.StopFrequency(n)
				} else {
					p.PlayFrequency(n)
				}
			}
		} else if midievent.IsNoteOff(midievent.Event(event.Status)) {
			n := 440 * math.Pow(2, (float64(event.Data1)-69)/12)
			for _, p := range mm.playables {
				p.StopFrequency(n)
			}
		}
	}

	return nil
}

func (mm *MidiMapper) Stop() error {
	err := mm.stream.Close()
	if err != nil {
		return err
	}

	return portmidi.Terminate()
}

func (mm *MidiMapper) Map(id int, p Playable) {
	mm.playables = append(mm.playables, p)
}
