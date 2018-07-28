package engine

type Track struct {
	ID               int
	Name             string
	InstrumentPatchs []*InstrumentPatch
}

func (t *Track) Next(i int) float64 {
	l := float64(len(t.InstrumentPatchs))
	var out float64

	for _, patch := range t.InstrumentPatchs {
		patch.Apply()
		out += patch.Next(i) / l
	}

	return out
}

func (t *Track) AddInstrumentPatch(patch *InstrumentPatch) {
	t.InstrumentPatchs = append(t.InstrumentPatchs, patch)
}

func NewTrack(name string) *Track {
	return &Track{
		ID:   -1,
		Name: name,
	}
}
