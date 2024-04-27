package main

import (
	"fmt"
	"os"
	"plugin"
	"strconv"
)

type OperatorPlugin interface {
	Operate(a, b float64) float64
}

func main() {
	// go run main.go addition 5 2
	//to test
	if len(os.Args) < 4 {
		fmt.Println("Usage: ./calculator [operator] [operand1] [operand2]")
		fmt.Println("operator can be : addition,division,subtract,multiply")
		os.Exit(1)
	}

	op := os.Args[1]
	operand1 := parseValues(os.Args[2])
	operand2 := parseValues(os.Args[3])

	operator := GetOperator(op)

	if operator != nil {
		result := operator.Operate(operand1, operand2)
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Invalid operator type")
	}
}

func parseValues(s string) float64 {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("Invalid number: %s\n", s)
		os.Exit(1)
	}
	return value
}

func GetOperator(op string) OperatorPlugin {

	var mod string = "./" + op + "/" + op + ".so"

	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	operatorX, err := plug.Lookup("Plugin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var operator OperatorPlugin
	operator, ok := operatorX.(OperatorPlugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}
	return operator

}
