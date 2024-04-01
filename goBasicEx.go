package main

import "fmt"

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven>OvenTime {
		return 0
	}
	return OvenTime-actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers*2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
return PreparationTime(numberOfLayers)+actualMinutesInOven 
}

func main() {
	fmt.Println("Remaining time in oven:", RemainingOvenTime(30))
	fmt.Println("Preparation time:", PreparationTime(2))
	fmt.Println("Elapsed time:", ElapsedTime(2, 30))
}
