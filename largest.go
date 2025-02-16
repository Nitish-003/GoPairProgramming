package main

import "fmt"

func main() {

	fmt.Println("Welcome")
	var first int
	var second int
	var third int

	fmt.Println("Enter the first Number")
	fmt.Scan(&first)
	fmt.Println("ENter the second Number")
	fmt.Scan(&second)
	fmt.Println("Enter the third Number")
	fmt.Scan(&third)

	if first > second && first > third {
		fmt.Println("First Number is the largest")
	} else if second > first && second > third {
		fmt.Println("Second Number is the largest")
	} else if third > second && second > first {
		fmt.Println("THord Number is the largest")
	} else {
		fmt.Println("All are equal")
	}

}
