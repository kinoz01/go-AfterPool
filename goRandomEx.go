package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(20)
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(12)) + rand.Float64()
	// return rand.Float64() * 12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	x := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
	return x
}

func main() {
	fmt.Println(RollADie())
	fmt.Println(GenerateWandEnergy())
	fmt.Println(ShuffleAnimals())
}
 
// Animal Magic
