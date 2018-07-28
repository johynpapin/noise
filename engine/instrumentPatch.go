package engine

import (
	"github.com/johynpapin/noise/things"
	"sync"
)

type patchNote struct {
	Hz         float64
	EntryPoint int
}

type InstrumentPatch struct {
	ID int

	Instrument *Instrument

	dropedNotes []*patchNote
	notes       []*patchNote

	lastEntryPoint int

	nextable *simpleNextable
	values   map[*things.Setting]float64
	lock     sync.RWMutex
}

type simpleNextable struct {
	Value float64
}

func newSimpleNextable() *simpleNextable {
	return &simpleNextable{}
}

func (sn *simpleNextable) Next(sample int, entryPoint int) float64 {
	return sn.Value
}

func NewInstrumentPatch(instrument *Instrument) *InstrumentPatch {
	return &InstrumentPatch{
		Instrument: instrument,
		nextable:   newSimpleNextable(),
		values:     make(map[*things.Setting]float64),
		lock:       sync.RWMutex{},
	}
}

func (ip *InstrumentPatch) Apply() {
	ip.Instrument.SetEntryPointNextable(ip.nextable)

	ip.lock.RLock()
	for _, setting := range ip.Instrument.Settings {
		setting.Value = ip.values[setting]
	}
	ip.lock.RUnlock()
}

func (ip *InstrumentPatch) Next(sample int) float64 {
	var o float64

	for _, note := range ip.notes {
		ip.nextable.Value = note.Hz
		o += ip.Instrument.Next(sample, note.EntryPoint) / 12
	}

	return o
}

func (ip *InstrumentPatch) Set(setting *things.Setting, value float64) {
	ip.lock.Lock()
	ip.values[setting] = value
	ip.lock.Unlock()
}

func (ip *InstrumentPatch) Get(setting *things.Setting) float64 {
	ip.lock.RLock()
	defer ip.lock.RUnlock()

	return ip.values[setting]
}

func (ip *InstrumentPatch) NoteOn(hz float64) {
	var entryPoint int

	if len(ip.notes) > 0 {
		entryPoint = ip.notes[len(ip.notes)-1].EntryPoint + 1
	} else if ip.lastEntryPoint != 0 {
		entryPoint = 0
	} else {
		entryPoint = ip.lastEntryPoint + 1
	}

	ip.lastEntryPoint = entryPoint

	ip.notes = append(ip.notes, &patchNote{
		Hz:         hz,
		EntryPoint: entryPoint,
	})

	if len(ip.notes) > 12 {
		ip.dropedNotes = append(ip.dropedNotes, ip.notes[0])
		ip.notes = ip.notes[1:]
	}
}

func (ip *InstrumentPatch) NoteOff(hz float64) {
	index := -1

	for i, n := range ip.dropedNotes {
		if n.Hz == hz {
			index = i
		}
	}

	if index > -1 {
		ip.dropedNotes = append(ip.dropedNotes[:index], ip.dropedNotes[index+1:]...)
		return
	}

	for i, n := range ip.notes {
		if n.Hz == hz {
			index = i
			break
		}
	}

	if index > -1 {
		ip.notes = append(ip.notes[:index], ip.notes[index+1:]...)
	}
}
