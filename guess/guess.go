package main

import (
	"bufio"
	"fmt"
	"os"
)

//TODOS
/*
1. print out how many tries it took
2. Identify when all numbers guessed, final number repeating
*/

//binary search log(n)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	low := 1
	high := 100
	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press enter when ready")
	scanner.Scan()
	guess_count := 0
	last_guess := 0

	for {
		//Binary search strategy
		guess_count++
		guess := (low + high) / 2
		if last_guess == guess && guess_count > 1 {
			fmt.Println("You're lying, I quit")
			break
		}
		last_guess = guess
		fmt.Println("I guess that the number is:", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) Too High?")
		fmt.Println("(b) Too Low?")
		fmt.Println("(c) Correct?")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won! It took me ", guess_count, "guesses.")
			break
		} else {
			fmt.Println("Invalid response, please try again")
		}
	}
}
