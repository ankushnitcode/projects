package main

import (
	"awesomeProject/Operators"
	"fmt"
	"os"
	"strconv"
)

type OperatorPlugin interface {
	Operate(a, b float64) float64
}

func main() {

	// go run main.go add 5 2
	//to test
	if len(os.Args) < 4 {
		fmt.Println("Usage: ./calculator [operator] [operand1] [operand2]")
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

func GetOperator(operatorType string) OperatorPlugin {
	switch operatorType {
	case "add":
		return Operators.AddPlugin{}
	case "subtract":
		return Operators.SubtractPlugin{}
	case "multiply":
		return Operators.MultiplyPlugin{}
	case "divide":
		return Operators.DividePlugin{}
	default:
		return nil
	}
}
