package main

import "fmt"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ten", "jack", "queen", "king": 
		return 10
	case "ace":
		return 11
	} 
	cards := []string{"two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, carte := range cards {
		if carte == card {
			return i+2
		}
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	a := ParseCard(card1)
	b := ParseCard(card2)
	c := ParseCard(dealerCard)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case a+b >= 21 && dealerCard != "ace" && c != 10:
		return "W"
	case (a+b >= 21 && (dealerCard == "ace" || c == 10)) || (a+b >= 17 && a+b <= 20) || (a+b >= 12 && a+b <= 16 && c < 7) :	
		return "S"
	}
	return "H"
}

func main() {
	value := ParseCard("two")
	fmt.Println(value)

	a := FirstTurn("ace", "ace", "jack") 
	b := FirstTurn("ace", "king", "ace") 
	c := FirstTurn("five", "queen", "ace")
	fmt.Printf("%s\n%s\n%s\n", a, b, c)
}

// Blackjack
