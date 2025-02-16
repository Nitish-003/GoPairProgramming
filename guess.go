package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var number int
	number = rand.Intn(1000)
	fmt.Println("Number is", number)

	for {
		var guess int
		fmt.Println("Guess the number (1-1000)")
		fmt.Scan(&guess)

		if number == guess {
			fmt.Println("You guessed the correct number")
			break
		} else if guess > number {
			fmt.Println("Number is greater than the original number! Try again")
		} else if guess < number {
			fmt.Println("Number is smaller than the original number! Try again")
		}

	}

}
