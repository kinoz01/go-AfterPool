package main

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func main() {
	var r rectangle
	r.length = 1.04
	r.width = 1.09

	rec:= rectangle{
		4.2,
		2.3,
	}
	rectangleInfo(rec)
	rectangleInfo(r)
}

func rectangleInfo(rec rectangle) {
	fmt.Println("Length:", rec.length)
	fmt.Println("width:", rec.width)
}

// struct practice
// https://headfirstgo.com/exercises/ch08.html
