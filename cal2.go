package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		// Display the menu
		fmt.Println("Calculator Menu:")
		fmt.Println("1: Addition")
		fmt.Println("2: Multiplication")
		fmt.Println("3: Power")
		fmt.Println("0: Exit")
		fmt.Print("Enter your choice: ")

		var value int
		fmt.Scan(&value)

		// Exit if the user chooses 0
		if value == 0 {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		}

		// Ask for two numbers
		var num1, num2 float64
		fmt.Scan(&num1, &num2)

		// Perform the chosen operation
		switch value {
		case 1:
			fmt.Printf("Result: %.2f + %.2f = %.2f\n", num1, num2, num1+num2)
		case 2:
			fmt.Printf("Result: %.2f * %.2f = %.2f\n", num1, num2, num1*num2)
		case 3:
			fmt.Printf("Result: %.2f ^ %.2f = %.2f\n", num1, num2, math.Pow(num1, num2))
		default:
			fmt.Println("Invalid choice! Please select 1, 2, 3, or 0 to exit.")
		}

		fmt.Println() // Add a blank line for readability
	}
}
