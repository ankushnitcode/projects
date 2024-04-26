package Operators

type MultiplyPlugin struct{}

func (p MultiplyPlugin) Operate(a, b float64) float64 {
	return a * b
}
