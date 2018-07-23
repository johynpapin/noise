package things

type Settings struct {
	Settings map[string]*Setting
}

func (s *Settings) add(thing Thing, name string) {
	s.Settings[name] = newSetting(name, thing)
}

func (s *Settings) Get(name string) float64 {
	return s.Settings[name].Get()
}

func (s *Settings) Set(name string, value float64) {
	s.Settings[name].Set(value)
}

func newSettings() *Settings {
	return &Settings{
		Settings: make(map[string]*Setting),
	}
}
