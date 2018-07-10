package main

type Audible interface {
	Next() float64
}

type Mutable interface {
	Mute()
}

type Sequencable interface {
	IsPlaying() bool
	Play()
	Stop()
}

type Playable interface {
	PlayFrequency(float64)
	StopFrequency(float64)
}
