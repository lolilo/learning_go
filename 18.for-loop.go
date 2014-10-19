package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)


	// As in C or Java, you can leave the pre and post statements empty.
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

}

