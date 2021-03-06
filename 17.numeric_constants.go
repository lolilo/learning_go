package main

import "fmt"

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { 
	return x*10 + 1 
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// fmt.Println(Big)
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))

	// overflow errors
	// In computer programming, an integer overflow occurs when an arithmetic operation 
	// attempts to create a numeric value that is too large to be represented 
	// within the available storage space. 
}
