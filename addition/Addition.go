package main

type AddPlugin struct{}

func (p AddPlugin) Operate(a, b float64) float64 {
	return a + b
}

var Plugin AddPlugin
