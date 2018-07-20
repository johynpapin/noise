package things

type IOS struct {
	inputs   *Inputs
	outputs  *Outputs
	settings *Settings
}

func NewIOS() *IOS {
	return &IOS{
		inputs:   newInputs(),
		outputs:  newOutputs(),
		settings: newSettings(),
	}
}

func (ios *IOS) GetInputs() *Inputs {
	return ios.inputs
}

func (ios *IOS) GetOutputs() *Outputs {
	return ios.outputs
}

func (ios *IOS) GetSettings() *Settings {
	return ios.settings
}
