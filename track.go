package main

type Track struct {
	audibles []Audible
}

func (t *Track) Next() []float32 {
	l := float32(len(t.audibles))
	out := make([]float32, 2)
	tmp := make([]float32, 2)

	for i := range t.audibles {
		tmp = t.audibles[i].Next()
		out[0] += tmp[0] / l
		out[1] += tmp[1] / l
	}

	return out
}

func (t *Track) AddAudible(a Audible) {
	t.audibles = append(t.audibles, a)
}

func NewTrack() *Track {
	return &Track{}
}
