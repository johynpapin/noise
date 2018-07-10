package main

type Track struct {
	audibles []Audible
}

func (t *Track) Next() float64 {
	l := float64(len(t.audibles))
	var out float64

	for i := range t.audibles {
		out += t.audibles[i].Next() / l
	}

	return out
}

func (t *Track) AddAudible(a Audible) {
	t.audibles = append(t.audibles, a)
}

func NewTrack() *Track {
	return &Track{}
}
