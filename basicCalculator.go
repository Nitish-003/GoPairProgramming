package main

import "fmt"

func main() {

	fmt.Println("Enter first number : ")
	var firstNumber float32
	fmt.Scan(&firstNumber)

	fmt.Println("Enter Second number : ")
	var secondNumber float32
	fmt.Scan(&secondNumber)

	fmt.Println("Enter the operation (+, -, *, /) : ")
	var operation string
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println("Sum is: ", firstNumber+secondNumber)
	case "-":
		fmt.Println("Difference is : ", firstNumber-secondNumber)
	case "*":
		fmt.Println("Multiply is : ", firstNumber*secondNumber)
	case "/":
		fmt.Println("DIvision is : ", firstNumber/secondNumber)
	default:
		fmt.Println("Invalid Choice")
	}

}
