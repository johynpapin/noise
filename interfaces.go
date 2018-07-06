package main

type Audible interface {
	Next() []float32
}

type Mutable interface {
	Mute()
}

type Sequencable interface {
	IsPlaying() bool
	Play()
	Stop()
}
