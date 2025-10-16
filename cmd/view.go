package main

import "fmt"

func MainMenu(round int) {
	fmt.Printf(`
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.

Round %d!

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)
`, round)
}

func selectDificulty(choice int) {
	var (
		e = "Easy"
		m = "Medium"
		h = "Hard"
	)
	switch choice {
	case 1:
		fmt.Printf("Great! You have selected the %s difficulty level.\nLet's start the game\n", e)
	case 2:
		fmt.Printf("Great! You have selected the %s difficulty level.\nLet's start the game\n", m)
	case 3:
		fmt.Printf("Great! You have selected the %s difficulty level.\nLet's start the game\n", h)
	}
}
