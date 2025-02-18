package main

import "fmt"

func main() {

	fmt.Println("Enter a Number : ")
	var n int
	fmt.Scan(&n)
	var sum int = 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Println("Sum of Even Numbers : ", sum)

}
