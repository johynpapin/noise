package things

type Setting struct {
	Name  string
	Value float64
	Thing Thing
}

func newSetting(name string, thing Thing) *Setting {
	return &Setting{
		Name:  name,
		Thing: thing,
	}
}

func (s *Setting) Set(value float64) {
	s.Value = value
}

func (s *Setting) Get() float64 {
	return s.Value
}
