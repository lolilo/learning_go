package main

import "fmt"

const Pi = 3.14

const (
	x = "this is an x"
	y = "this is a y"
)

func main() {
	// Go can read Chinese. Awesome. :D 
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	fmt.Printf("%T\n", Pi)


	fmt.Println(x, "\n", y)
	const Truth = true
	fmt.Println("Go rules?", Truth)
}