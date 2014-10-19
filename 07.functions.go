// http://tour.golang.org/#10

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// this is pretty interesting
// Functions take parameters. 
// In Go, functions can return multiple "result parameters", not just a single value. 
// They can be named and act just like variables.

// If the result parameters are named, 
// a return statement without arguments returns the current values of the results.

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
