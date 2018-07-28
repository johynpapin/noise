package things

type Nextable interface {
	Next(sample int, entryPoint int) float64
}

type Output struct {
	ID             int
	Name           string
	Nextable       Nextable
	Thing          *Thing
	current        float64
	lastSample     int
	lastEntryPoint int
}

func newOutput(name string, thing *Thing, nextable Nextable) *Output {
	return &Output{
		Name:     name,
		Nextable: nextable,
		Thing:    thing,
	}
}

func NewEntryPoint(nextable Nextable) *Output {
	return &Output{
		Name:           "Entry Point",
		Nextable:       nextable,
		lastSample:     -1,
		lastEntryPoint: -1,
	}
}

func (o *Output) Next(sample int, entryPoint int) float64 {
	if sample != o.lastSample || entryPoint != o.lastEntryPoint {
		o.current = o.Nextable.Next(sample, entryPoint)
	}

	o.lastSample = sample
	o.lastEntryPoint = entryPoint

	return o.current
}
