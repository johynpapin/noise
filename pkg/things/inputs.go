package things

type Inputs struct {
	Inputs map[string]*Input
}

func (i *Inputs) Read(name string) float64 {
	input, _ := i.Inputs[name]
	return input.Read()
}

func (i *Inputs) add(thing Thing, name string, defaultValue float64) {
	i.Inputs[name] = newInput(name, thing, defaultValue)
}

func (i *Inputs) Attach(inputName string, output *Output) {
	input, _ := i.Inputs[inputName]
	input.Attach(output)
}

func newInputs() *Inputs {
	return &Inputs{
		Inputs: make(map[string]*Input),
	}
}
