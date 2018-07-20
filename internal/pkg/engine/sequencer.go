package engine

import (
	"github.com/johynpapin/noise/pkg/constants"
	log "github.com/sirupsen/logrus"
)

type Sequencer struct {
	timer        *Timer
	sequencables []Sequencable
	pattern      *Pattern
}

func NewSequencer(timer *Timer, pattern *Pattern) *Sequencer {
	return &Sequencer{timer, nil, pattern}
}

func (s *Sequencer) Start() {
	ppqnCount := 0
	beat := 0

	for {
		select {
		case <-s.timer.Pulses:
			ppqnCount++

			if ppqnCount%(int(constants.PPQN)/4) == 0 {
				go s.nextBeat(beat)
				beat++
			}

			if ppqnCount == (int(constants.PPQN) * 4) {
				ppqnCount = 0
				beat = 0
			}
		}
	}
}

func (s *Sequencer) nextBeat(beat int) {
	for i := range s.sequencables {
		if s.pattern.AtBeat(beat) {
			log.Info("woaw!")
			if !s.sequencables[i].IsPlaying() {
				s.sequencables[i].Play()
			}
		} else {
			s.sequencables[i].Stop()
		}
	}
}

func (s *Sequencer) AddSequencable(sequencable Sequencable) {
	s.sequencables = append(s.sequencables, sequencable)
}
