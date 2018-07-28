package engine

import (
	log "github.com/sirupsen/logrus"
)

type Sequencer struct {
	clock       *Clock
	midiManager *MidiManager

	loops []*Loop
}

func NewSequencer(clock *Clock, midiManager *MidiManager) *Sequencer {
	return &Sequencer{
		clock:       clock,
		midiManager: midiManager,
	}
}

func (s *Sequencer) AddLoop(loop *Loop) {
	s.loops = append(s.loops, loop)
}

func (s *Sequencer) Start() {
	go func() {
		for {
			<-s.clock.Pulses
			log.Info("pulse")

			for _, loop := range s.loops {
				if loop.IsRecording {
					log.Info("recording ! pulse ++")
					loop.Pulses++
				} else {
					log.Info("playing ! pulse ++")
					loop.Pulses++
					if loop.TotalPulses < loop.Pulses {
						log.Info("end ! pulse = 0")
						loop.Pulses = 0
					}

					messages := loop.GetMessages(loop.Pulses)

					for _, msg := range messages {
						log.Info("olala message ! olala message !")
						s.midiManager.ProcessLoopMessage(loop.InstrumentPatch, msg)
					}
				}
			}
		}
	}()
}
