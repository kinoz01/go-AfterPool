package main

import "fmt"

// Declaring a Struct
type Animal struct {
	name   string
	weight int
}

func main() {
	// Creating an instance of a struct
	var dog Animal
	dog.name = "Dog"
	dog.weight = 40
	fmt.Println(dog)

	// Creating an instance using struct literate
	var cat = Animal{name: "Cat", weight: 5}
	fmt.Println(cat)

	// Creating an instance using the new keyword
	var bird = new(Animal)
	bird.name = "Bird"
	bird.weight = 1
	fmt.Println(bird)

	// Creating an instance using the pointer address operator
	var monkey = &Animal{name: "Monkey", weight: 10}
	fmt.Println(monkey)

	// Creating an instance using :=
	lion := Animal{
		"lion", // this is shortcut usually u have  
		230,	// to write name = "lion"
	}
	fmt.Println(lion)

	// Comparing struct instances
	fmt.Println(bird == monkey)

	// Copying struct type using pointer reference
	monkey2 := monkey
	monkey2.name = "Monkey2"
	fmt.Println(monkey2)
}

// struct practice
// https://github.com/TannerGabriel/learning-go/blob/master/basics/10-Struct/Struct.go
