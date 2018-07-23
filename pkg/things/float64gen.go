package things

func init() {
	registerThing("Float64Gen", NewFloat64Gen)
}

type Float64Gen struct {
	*IOS
}

func NewFloat64Gen() Thing {
	t := &Float64Gen{
		IOS: NewIOS(),
	}

	t.IOS.settings.add(t, "value")

	t.IOS.outputs.add(t, "output", NewFloat64GenOutput(t))

	return t
}

// Float64GenOutput

type Float64GenOutput struct {
	t *Float64Gen
}

func NewFloat64GenOutput(t *Float64Gen) *Float64GenOutput {
	return &Float64GenOutput{
		t: t,
	}
}

func (o Float64GenOutput) Next() float64 {
	return o.t.settings.Get("value")
}
