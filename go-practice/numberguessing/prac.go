package main

import "fmt"

func main() {
	number := 50
	guess := 105
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 to 100")
	}
	if guess >= 1 && guess <= 100 {

		if guess < number {
			fmt.Println("Too Low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("You Got it")
		}
	}
	fmt.Println(number <= guess, number >= guess, number != guess)
}
