package main

import (
	"math"
)

type StereoSine struct {
	stepL, phaseL float64
	stepR, phaseR float64
	playing       bool
}

func (ss *StereoSine) IsPlaying() bool {
	return ss.playing
}

func (ss *StereoSine) Play() {
	ss.playing = true
}

func (ss *StereoSine) Stop() {
	ss.playing = false
	ss.phaseL, ss.phaseR = 0, 0
}

func (ss *StereoSine) Next() []float32 {
	if !ss.playing {
		return []float32{0, 0}
	}

	out := make([]float32, 2)

	out[0] = float32(math.Sin(2 * math.Pi * ss.phaseL))
	_, ss.phaseL = math.Modf(ss.phaseL + ss.stepL)

	out[1] = float32(math.Sin(2 * math.Pi * ss.phaseR))
	_, ss.phaseR = math.Modf(ss.phaseR + ss.stepR)

	return out
}

func NewStereoSine(freqL, freqR, sampleRate float64) *StereoSine {
	return &StereoSine{freqL / sampleRate, 0, freqR / sampleRate, 0, false}
}
