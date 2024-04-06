package main

import (
	"fmt"
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	c := 0
	if n <= 0 {
		// err := fmt.Errorf("Integers must be greater than zero")
		// return 0, err
        return 0, errors.New("input must be a positive integer")
    }
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		c++
	}
	return c, nil
}

func main() {
	fmt.Println(CollatzConjecture(12))
}

// Collatz Conjecture
