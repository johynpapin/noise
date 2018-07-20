package things

type Thing interface {
	GetInputs() *Inputs
	GetOutputs() *Outputs
	GetSettings() *Settings
}
