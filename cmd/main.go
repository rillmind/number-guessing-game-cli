package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	var round int
	for {
		var choice, number, attempts int

		round++

		MainMenu(round)

		fmt.Print("\nYour choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			selectDificulty(choice)

			start := time.Now()

			rNumber := rand.IntN(101)

			for range 10 {
				fmt.Print("\nYour guess: ")
				fmt.Scan(&number)
				fmt.Println()

				if number < rNumber {
					fmt.Printf("Incorrent! The number is greater than %d\n", number)
				} else if number > rNumber {
					fmt.Printf("Incorrent! The number is less than %d\n", number)
				} else if number == rNumber {
					fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
					break
				}

				attempts++
			}

			fmt.Printf("\nTime passed: %s\n", time.Since(start))

		case 2:
			selectDificulty(choice)

			start := time.Now()

			rNumber := rand.IntN(101)

			for range 5 {
				fmt.Print("\nYour guess: ")
				fmt.Scan(&number)
				fmt.Println()

				if number < rNumber {
					fmt.Printf("Incorrent! The number is greater than %d\n", number)
				} else if number > rNumber {
					fmt.Printf("Incorrent! The number is less than %d\n", number)
				} else if number == rNumber {
					fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
					break
				}

				attempts++
			}

			fmt.Printf("\nTime passed: %s\n", time.Since(start))

		case 3:
			selectDificulty(choice)

			start := time.Now()

			rNumber := rand.IntN(101)

			for range 3 {
				fmt.Print("\nYour guess: ")
				fmt.Scan(&number)
				fmt.Println()

				if number < rNumber {
					fmt.Printf("Incorrent! The number is greater than %d\n", number)
				} else if number > rNumber {
					fmt.Printf("Incorrent! The number is less than %d\n", number)
				} else if number == rNumber {
					fmt.Printf("Congratulations! You guessed the correct number in %d attempts\n", attempts)
					break
				}

				attempts++
			}

			fmt.Printf("\nTime passed: %s\n", time.Since(start))
		}
	}
}
