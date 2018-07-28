package things

type Settings struct {
	Settings map[string]*Setting
}

func (s *Settings) add(thing *Thing, name string, defaultValue float64, min float64, max float64) {
	s.Settings[name] = newSetting(name, thing, defaultValue, min, max)
}

func (s *Settings) Get(name string) *Setting {
	return s.Settings[name]
}

func newSettings() *Settings {
	return &Settings{
		Settings: make(map[string]*Setting),
	}
}
