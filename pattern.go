package main

type Pattern struct {
	pattern []bool
}

func NewPattern(pattern []bool) *Pattern {
	return &Pattern{pattern}
}

func (p *Pattern) AtBeat(beat int) bool {
	return p.pattern[beat]
}
