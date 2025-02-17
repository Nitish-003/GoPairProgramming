package main

import "fmt"

func main() {

	var number int
	fmt.Println("Enter the number ")
	fmt.Scan(&number)

	var binary []int

	for number != 0 {
		var a = number % 2
		binary = append([]int{a}, binary...)
		number = number / 2
	}

	var binaryStr string

	for _, bit := range binary {
		binaryStr = binaryStr + fmt.Sprintf("%d", bit)
	}

	fmt.Println("Binary conversion is ", binaryStr)

}
