package things

type Thing struct {
	ID       int
	Name     string
	Inputs   *Inputs
	Outputs  *Outputs
	Settings *Settings
}

func newThing(name string) *Thing {
	return &Thing{
		Name:     name,
		Inputs:   newInputs(),
		Outputs:  newOutputs(),
		Settings: newSettings(),
	}
}
