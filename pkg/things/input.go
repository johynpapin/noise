package things

type Input struct {
	Name         string
	Output       *Output
	Thing        Thing
	DefaultValue float64
}

func newInput(name string, thing Thing, defaultValue float64) *Input {
	return &Input{
		Name:         name,
		Thing:        thing,
		DefaultValue: defaultValue,
	}
}

func (input *Input) Attach(output *Output) {
	input.Output = output
}

func (input *Input) Detach() {
	input.Output = nil
}

func (input *Input) Read() float64 {
	if input.Output == nil {
		return input.DefaultValue
	}

	return input.Output.Next()
}
