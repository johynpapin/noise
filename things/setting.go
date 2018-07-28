package things

type Setting struct {
	ID          int
	Name        string
	ExposedName string
	Value       float64
	Min         float64
	Max         float64
	Thing       *Thing
}

func newSetting(name string, thing *Thing, defaultValue float64, min float64, max float64) *Setting {
	return &Setting{
		Name:  name,
		Value: defaultValue,
		Thing: thing,
		Min:   min,
		Max:   max,
	}
}
