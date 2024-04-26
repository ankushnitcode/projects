package Operators

type DividePlugin struct{}

func (p DividePlugin) Operate(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
