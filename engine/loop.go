package engine

import (
	"github.com/gomidi/midi"
)

type Loop struct {
	ID int

	messages    map[int][]midi.Message
	Pulses      int
	TotalPulses int
	IsRecording bool
	IsFinished  bool

	InstrumentPatch *InstrumentPatch
}

func NewLoop(patch *InstrumentPatch) *Loop {
	return &Loop{
		messages:        make(map[int][]midi.Message),
		InstrumentPatch: patch,
	}
}

func (l *Loop) SaveMessage(message midi.Message) {
	l.messages[l.Pulses] = append(l.messages[l.Pulses], message)
}

func (l *Loop) GetMessages(pulse int) []midi.Message {
	return l.messages[pulse]
}

func (l *Loop) StartRecording() {
	l.IsRecording = true
	l.Pulses = 0
}

func (l *Loop) StopRecording() {
	l.IsRecording = false
	l.TotalPulses = l.Pulses
	l.Pulses = 0
}
