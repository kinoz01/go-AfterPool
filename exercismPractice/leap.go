package main

import "fmt"


func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%100 == 0 && year%400 == 0)
}

func main() {
	fmt.Println(IsLeapYear(1997))
	fmt.Println(IsLeapYear(1900))
	fmt.Println(IsLeapYear(2000))
}

// Leap
