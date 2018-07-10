package stuff

type Float64Generator struct {
	Value float64
}

func NewFloat64Generator(value float64) *Float64Generator {
	return &Float64Generator{
		Value: value,
	}
}

func (s *Float64Generator) Next() float64 {
	return s.Value
}
