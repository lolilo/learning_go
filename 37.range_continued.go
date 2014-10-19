package main

import "fmt"

func main() {
	pow := make([]int, 10)
	
	// if you only want the index, dro pthe ", value" entirely
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// skip the index or value by assigning to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}



// uint is an unsigned integer type that is at least 32 bits in size. 
	// It is a distinct type, however, and not an alias for, say, uint32.

// uint16 is the set of all unsigned 16-bit integers. Range: 0 through 65535.
// uint32 is the set of all unsigned 32-bit integers. Range: 0 through 4294967295.
