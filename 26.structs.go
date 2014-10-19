// A struct is a collection of fields.

// (And a type declaration does what you'd expect.)

// Struct fields are accessed using a dot.

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)

	fmt.Println(Vertex{1, 2})
}
