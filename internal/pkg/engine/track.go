package engine

import (
	"github.com/johynpapin/noise/pkg/things"
)

type Track struct {
	Name          string           `json:"name"`
	Outputs       []*things.Output `json:"outputs"`
	ThingsManager *things.ThingsManager
	Things        []things.Thing
}

func (t *Track) Next() float64 {
	l := float64(len(t.Outputs))
	var out float64

	for _, output := range t.Outputs {
		out += output.Next() / l
	}

	return out
}

func (t *Track) AddOutput(output *things.Output) {
	t.Outputs = append(t.Outputs, output)
}

func (t *Track) AddThing(thing things.Thing) {
	t.Things = append(t.Things, thing)
}

func newTrack(name string, thingsManager *things.ThingsManager) *Track {
	return &Track{
		Name:          name,
		ThingsManager: thingsManager,
	}
}
