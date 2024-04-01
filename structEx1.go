package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func main() {
	var r rectangle
	r.length = 2
	r.width = 4

	rec:= rectangle{
		4.2,
		2.3,
	}
	rectangleInfo(rec)
	rectangleInfo(r)
	makeSquare(&rec) // Pass a pointer to makeSquare.
	makeSquare(&r)   // Pass a pointer to makeSquare.
	rectangleInfo(rec)
	rectangleInfo(r)
}

func rectangleInfo(rec rectangle) {
	fmt.Println("Length:", rec.length)
	fmt.Println("width:", rec.width)
}

// Modifying Structs from Functions :

/* Accepts a pointer to a rectangle rather than a
rectangle value, so that it can modify the original
value at the pointer.
*/
func makeSquare(rec *rectangle) {
	if rec.length > rec.width {
		rec.length = rec.width
	} else {
		rec.width = rec.length
	}
	// Remember that the dot operator works the same
	// with a pointer to a struct as it does with the
	// actual struct. You don't have to explicitly
	// write (*r).length or (*r).width.
}

// struct practice
// https://headfirstgo.com/exercises/ch08.html
