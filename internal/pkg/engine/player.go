package engine

import (
	"github.com/gordonklaus/portaudio"
	"github.com/johynpapin/noise/pkg/things"
)

type Player struct {
	stream *portaudio.Stream

	Tracks        map[string]*Track
	ThingsManager *things.ThingsManager
}

func NewPlayer() *Player {
	return &Player{
		Tracks:        make(map[string]*Track),
		ThingsManager: things.NewThingsManager(),
	}
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
			tmp = float32(t.Next())
			out[0][i] += tmp / l
			out[1][i] += tmp / l
		}
	}
}

func (p *Player) NewTrack(name string) {
	p.Tracks[name] = newTrack(name, p.ThingsManager)
}

func (p *Player) GetTrack(name string) *Track {
	return p.Tracks[name]
}
