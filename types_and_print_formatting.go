package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe 		bool	 	= false
	MaxInt 		uint64 		= 1<<64 - 1
	z 			complex128 	= cmplx.Sqrt(-5 + 12i)
)


func main() {
	const f = "%T(%v)\n" 

	// http://golang.org/pkg/fmt/
	// Printf means to "print formatted"
	// %T	a Go-syntax representation of the type of the value
	// %v	the value in a default format

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
