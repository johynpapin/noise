package things

type State struct {
	used              []int
	currentSample     int
	currentEntryPoint int
	states            map[int]map[string]interface{}
}

func (s *State) clean() {
	var used bool

	for entryPoint := range s.states {
		used = false
		for _, ep := range s.used {
			if entryPoint == ep {
				used = true
				break
			}
		}

		if !used {
			delete(s.states, entryPoint)
		}
	}

	s.used = nil
}

func (s *State) get(name string, sample int, entryPoint int) interface{} {
	if sample != s.currentSample {
		s.clean()
		s.currentSample = sample
	}

	s.currentEntryPoint = entryPoint
	s.used = append(s.used, entryPoint)

	value, ok := s.states[entryPoint]
	if !ok {
		s.states[entryPoint] = make(map[string]interface{})
		return nil
	}

	return value[name]
}

func (s *State) GetFloat64(name string, sample int, entryPoint int) float64 {
	value, _ := s.get(name, sample, entryPoint).(float64)
	return value
}

func (s *State) GetInt(name string, sample int, entryPoint int) int {
	value, _ := s.get(name, sample, entryPoint).(int)
	return value
}

func (s *State) GetADSRState(name string, sample int, entryPoint int) ADSRState {
	value, _ := s.get(name, sample, entryPoint).(ADSRState)
	return value
}

func (s *State) Set(name string, value interface{}) {
	s.states[s.currentEntryPoint][name] = value
}

func newState() *State {
	return &State{
		currentSample:     -1,
		currentEntryPoint: -1,
		states:            make(map[int]map[string]interface{}),
	}
}
