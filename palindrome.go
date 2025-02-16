package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numberInput string
	fmt.Println("Enter a number")
	fmt.Scan(&numberInput)

	var number int64
	var err error
	number, err = strconv.ParseInt(numberInput, 10, 64)
	if err != nil {

		temp := number
		var n int64

		for number != 0 {
			var a = number % 10
			n = n*10 + a
			number = number / 10
		}

		if temp == n {
			fmt.Println("It is a palindrome number")
		} else {
			fmt.Println("It is not a palindrome number")
		}
	} else {
		fmt.Println("Invalid input")
	}

}
