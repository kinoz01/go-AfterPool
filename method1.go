package main

import "fmt"

/*
in structEx1.go we had you write a rectangleInfo function 
that accepted a rectangle struct value and printed its 
length and width fields. Convert the rectangleInfo 
function into an info method on the rectangle type.
*/

type rectangle struct {
	length float64
	width  float64
}

// converting this func :
func rectangleInfo(r rectangle) {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}
// to a method : 
// Convert the regular parameter to a receiver parameter
// on the rectangle type.
func (r *rectangle) info() {
	fmt.Println("Length:", r.length)
	fmt.Println("Width:", r.width)
}

func main() {
	var r rectangle
	r.length = 4.2
	r.width = 2.3
	//rectangleInfo(r)
	// Changing the function call to a method call.
	r.info()
	r.makeSquare()
	r.info()
}

// Convert this function to a method on the "rectangle" type.
func makeSquare(r *rectangle) {
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}
func (r *rectangle) makeSquare() {
	// Again, no need to write (*r).length; Go gets
	// the value at the pointer automatically.
	if r.length > r.width {
		r.length = r.width
	} else {
		r.width = r.length
	}
}

// https://headfirstgo.com/exercises/ch09.html
