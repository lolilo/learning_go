// Go has pointers, but no pointer arithmetic.

// Struct fields can be accessed through a struct pointer. 
// The indirection through the pointer is transparent.

// Google "pointer in Go"
// http://www.golang-book.com/8/index.htm

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p)
}
