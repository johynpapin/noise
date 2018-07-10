package stuff

type Input struct {
	Stuff Stuff
}

func NewInput() *Input {
	return &Input{}
}

func (input *Input) Attach(stuff Stuff) {
	input.Stuff = stuff
}

func (input *Input) Next() float64 {
	if input.Stuff == nil {
		return 0
	}

	return input.Stuff.Next()
}
