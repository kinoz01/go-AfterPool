package main

import "fmt"

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		return len(layers)*2
	}
	return len(layers)*prepTime
}

func Quantities(layers []string) (int, float64) {
	countSauce  := 0.0
	countNoodle := 0
	for _, ingredient := range layers {
		if ingredient == "sauce" {
			countSauce++
		}
		if ingredient == "noodles" {
			countNoodle++
		}
	}
	return 50*countNoodle, 0.2*countSauce
} 

func AddSecretIngredient(friendRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendRecipe[len(friendRecipe)-1]
}

func ScaleRecipe(PortionsforTwo []float64, portion int) []float64 {
	scaledPortions := []float64{}
	for _, IngredientPortion := range PortionsforTwo {
		scaledPortions = append(scaledPortions, IngredientPortion*float64(portion)/2)
	}
	return scaledPortions
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(PreparationTime(layers, 0))
	
	fmt.Println(Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}))
	// => 100, 0.4

	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}
	AddSecretIngredient(friendsList, myList)
	fmt.Printf("%#v\n", myList)

	quantities := []float64{ 1.2, 3.6, 10.5 }
	scaledQuantities := ScaleRecipe(quantities, 4)
	fmt.Printf("%#v\n", scaledQuantities)
}

// Lasagna Master
