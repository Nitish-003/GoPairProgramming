package main

import "fmt"

func main() {

	var number int
	fmt.Println("Enter the number")
	fmt.Scan(&number)

	result := factorial(number)
	fmt.Println("Result is ", result)

}

func factorial(number int) int {

	if number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}

}
