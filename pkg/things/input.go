package things

type Input struct {
	Name   string
	Output *Output
	Thing  Thing
}

func newInput(name string, thing Thing) *Input {
	return &Input{
		Name:  name,
		Thing: thing,
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
		return 0
	}

	return input.Output.Next()
}
