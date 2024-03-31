package main

import "fmt"

type Shape struct {
	name string
	size int
}

func main() {

	/*
	   In Go, a struct is a sequence of named elements
	   called fields, each field has a name and type.
	   The name of a field must be unique within the struct.
	   Structs can be compared with classes in the
	   Object-Oriented Programming paradigm.
	*/

	/*
	   Field names in structs follow the Go convention -
	   fields whose name starts with a lower case letter
	   are only visible to code in the same package, whereas
	   those whose name starts with an upper case letter are
	   visible in other packages.
	*/

	// Creating instances of a struct

	s := Shape{
		name: "Square",
		size: 25,
	}

	// To read or modify instance fields, use the . notation:

	s.name = "Circle"
	s.size = 35
	fmt.Printf("name: %s, size: %d\n", s.name, s.size)

	/* Fields that don't have an initial value assigned,
	   will have their zero value. For example:
	*/

	c := Shape{name: "Circle"}
	fmt.Printf("name: %s, size: %d\n", c.name, c.size)

	fmt.Print("first struct : ")
	fmt.Println(s)
	fmt.Print("second struct : ")
	fmt.Println(c)

	/*
	   In the following example, one of these New functions
	   is used to create a new instance of Shape
	   and automatically set a default value for the size
	   of the shape:
	*/

	fmt.Print("new struct using NewShape function : ")
	fmt.Println(NewShape("Triangle"))

	/*
	   Another unusual way of creating a new instance of a struct
	   is by using the new built-in:
	*/

	empty := new(Shape) // s will be of type *Shape (pointer to shape)
	fmt.Printf("name: %s size: %d\n", empty.name, empty.size)
}

// "New" functions

/* Sometimes it can be nice to have functions that
help us create struct instances. By convention,
these functions are usually called New or have their
names starting with New, but since they are just
regular functions, you can give them any name you want.
They might remind you of constructors in other
languages, but in Go they are just regular
functions :
*/

func NewShape(s string) Shape {
	return Shape{
		name: s,
		size: 100, //default-value for size is 100
	}
}

// https://exercism.org/tracks/go/concepts/structs
