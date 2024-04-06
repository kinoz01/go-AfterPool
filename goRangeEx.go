package main

import "fmt"

type File []bool

type Chessboard map[string]File
	
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	c := 0
	for key, value := range cb {
		if key == file {
			for _, vraiFaux := range value {
				if vraiFaux {
					c++
				}
			}
		}
	}	
	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	c := 0
	if rank > 8 || rank < 1 {
		return c
	}
	for _, value := range cb {
		for i, vraiFaux := range value {
			if vraiFaux && i+1 == rank {
				c++
			}
		}
	}	
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	c := 0
	for _, value := range cb {
		// c += len(value)
		for i := 0; i < len(value); i++ {
			c++
		}
	}
	return c
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	c := 0
	for _, value := range cb {
		for _, vraiFaux := range value {
			if vraiFaux {
				c++
			}
		}
	}	
	return c
}

func main() {
	board := Chessboard{
		"A": {true, false, true, false, false, false, false, true},
    	"B": {false, false, false, false, true, false, false, false},
		"C": {false, false, true, false, false, false, false, false},
		"D": {false, false, false, false, false, false, false, false},
		"E": {false, false, false, false, false, true, false, true},
		"F": {false, false, false, false, false, false, false, false},
		"G": {false, false, false, true, false, false, false, false},
		"H": {true, true, true, true, true, true, false, true},
	}
	fmt.Println(CountInFile(board, "A"))
	fmt.Println(CountInRank(board, 2))
	fmt.Println(CountAll(board))
	fmt.Println(CountOccupied(board))
}

// Chessboard
