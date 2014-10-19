package main

import "fmt"

// var i int
// var c, python, java bool


// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; 
// the variable will take the type of the initializer.

var i, j int = 1, 2
var c, python, java = true, false, "no!"


func main() {
	fmt.Println(i, j, c, python, java)
}