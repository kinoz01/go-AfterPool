package main

import "fmt"

func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}

func main() {
	fmt.Println(ShareWith("Ayoub"))
}

// Two Fer: if condition / basics
