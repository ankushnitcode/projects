package main

type SubtractPlugin struct{}

func (p SubtractPlugin) Operate(a, b float64) float64 {
	return a - b
}

var Plugin SubtractPlugin
