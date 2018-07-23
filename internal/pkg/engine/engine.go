package engine

type Engine struct {
	Player      *Player
	MidiManager *MidiManager
}

func NewEngine() *Engine {
	return &Engine{
		Player:      NewPlayer(),
		MidiManager: NewMidiManager(),
	}
}

func (e *Engine) Start() error {
	err := e.Player.Start()
	if err != nil {
		return err
	}

	err = e.MidiManager.Start()
	if err != nil {
		return err
	}

	return nil
}
