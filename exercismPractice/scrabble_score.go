package main

import (
	"fmt"
	"strings"
)

func Score(word string) int {
	c := 0
	for _, char := range strings.ToUpper(word) {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S','T':
			c += 1
		case 'D', 'G':
			c += 2
		case 'B', 'C', 'M', 'P': 
			c += 3
		case 'F', 'H', 'V', 'W', 'Y':
			c += 4
		case 'K':
			c += 5
		case 'J', 'X':
			c += 8
		case 'Q', 'Z':
			c += 10
		}
	}
	return c
}

func main() {
	fmt.Println(Score("cabbage"))
}

// Scrabble Score
