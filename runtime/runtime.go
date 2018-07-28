package runtime

import (
	"github.com/johynpapin/noise/engine"
	"github.com/johynpapin/noise/store"
)

type Runtime struct {
	Player      *engine.Player
	MidiManager *engine.MidiManager
	Store       *store.Store
	Clock       *engine.Clock
	Sequencer   *engine.Sequencer
}

func NewRuntime(store *store.Store) *Runtime {
	r := &Runtime{
		Player:      engine.NewPlayer(),
		MidiManager: engine.NewMidiManager(),
		Store:       store,
		Clock:       engine.NewClock(120),
	}

	r.Sequencer = engine.NewSequencer(r.Clock, r.MidiManager)

	return r
}

func (e *Runtime) Start() error {
	err := e.Player.Start()
	if err != nil {
		return err
	}

	err = e.MidiManager.Start()
	if err != nil {
		return err
	}

	e.Clock.Start()
	e.Sequencer.Start()

	return nil
}
