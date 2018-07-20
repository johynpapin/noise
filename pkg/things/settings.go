package things

type Settings struct {
	settings map[string]interface{}
}

func (s *Settings) GetFloat64(name string) float64 {
	var value float64

	rawValue, ok := s.settings[name]
	if !ok {
		return value
	}

	value, ok = rawValue.(float64)

	return value
}

func (s *Settings) Set(name string, value interface{}) {
	s.settings[name] = value
}

func newSettings() *Settings {
	return &Settings{
		settings: make(map[string]interface{}),
	}
}
