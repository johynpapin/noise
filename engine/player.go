package engine

import (
	"github.com/gordonklaus/portaudio"
)

type Player struct {
	stream *portaudio.Stream

	Tracks []*Track
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Start() error {
	err := portaudio.Initialize()
	if err != nil {
		return err
	}

	p.stream, err = portaudio.OpenDefaultStream(0, 2, 44100, 0, p.processAudio)
	if err != nil {
		return err
	}

	return p.stream.Start()
}

func (p *Player) Stop() error {
	err := p.stream.Stop()
	if err != nil {
		return err
	}

	return portaudio.Terminate()
}

func (p *Player) processAudio(out [][]float32) {
	l := float32(len(p.Tracks))
	var tmp float32

	for i := range out[0] {
		out[0][i], out[1][i] = 0, 0
		for _, t := range p.Tracks {
			tmp = float32(t.Next(i))
			out[0][i] += tmp / l
			out[1][i] += tmp / l
		}
	}
}

func (p *Player) AddTrack(track *Track) {
	p.Tracks = append(p.Tracks, track)
}
