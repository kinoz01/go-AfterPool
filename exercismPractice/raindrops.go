package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	var r string
	if number%3 == 0 {
		r += "Pling"
	}
	if number%5 == 0 {
		r += "Plang"
	}
	if number%7 == 0 {
		r += "Plong"
	} 
	if r == "" {
		r = strconv.Itoa(number)
	}
	return r
}

func main() {
	fmt.Println(Convert(28))
	fmt.Println(Convert(30))
	fmt.Println(Convert(34))
}


// Raindrops / if statements
