package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)	
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", " "))
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	fmt.Println(AddBorder("Welcome!", 10))

	message := `**************************
*    BUY NOW, SAVE 10%   *
**************************
`	
	//fmt.Println(message)
	fmt.Println(CleanupMessage(message))
}
