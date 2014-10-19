package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	ans := x / 2
	for i := 0; i < 10000; i++ {
		ans = ans - (ans*ans - x) / 2 * ans
	}
	return ans
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
