package things

type Outputs struct {
	Outputs map[string]*Output
}

func newOutputs() *Outputs {
	return &Outputs{
		Outputs: make(map[string]*Output),
	}
}

func (o *Outputs) add(thing Thing, name string, nextable nextable) {
	o.Outputs[name] = newOutput(name, thing, nextable)
}

func (o *Outputs) Get(name string) *Output {
	return o.Outputs[name]
}
