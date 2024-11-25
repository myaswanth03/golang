package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string

	fmt.Println("enter the first number : ")
	fmt.Scanln(&number1)

	fmt.Println("enter the second number : ")
	fmt.Scanln(&number2)

	fmt.Println("enter the operator : ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1*number2)

	case "/":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1/number2)

	default:
		fmt.Println("Invalid operator")
	}
}
