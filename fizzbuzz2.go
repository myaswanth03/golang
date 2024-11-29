package main

import "fmt"

func main() {
	fmt.Println("enter integer : ")
	var input int
	fmt.Scan(&input)

	for i := 1; i <= input; i++ {

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "buzz")
		}
	}
}
