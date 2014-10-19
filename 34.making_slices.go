// Gets or sets the total number of elements the internal data structure can hold 
// without resizing.

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2] // this is interesting
	printSlice("c", c)
	d := c[2:5] // this is interesting
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", 
		s, len(x), cap(x), x)
}
