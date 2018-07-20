package things

type nextable interface {
	Next() float64
}

type Output struct {
	Name     string
	nextable nextable
	Thing    Thing
}

func newOutput(name string, thing Thing, nextable nextable) *Output {
	return &Output{
		Name:     name,
		nextable: nextable,
		Thing:    thing,
	}
}

func (o *Output) Next() float64 {
	return o.nextable.Next()
}
