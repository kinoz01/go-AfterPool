package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
	func FavoriteCards() []int {
		return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) || index < 0 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}	
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	for _, n := range slice {
		values = append(values, n)
	}
	return values
	// return append(values,slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return slice
	}
	r := []int{}
	for i, n := range slice {
		if i != index {
			r = append(r, n)
		}
	}
	return r
	// return append(slice[:index], slice[index+1:]...)
}

func main() {
	fmt.Println(FavoriteCards())

	card := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card)

	index := 3
	newCard := 6
	cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)

	slice := []int{3, 2, 6, 4, 8}
	cardss := PrependItems(slice, 5, 1)
	fmt.Println(cardss)

	cardsss := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cardsss)
}
