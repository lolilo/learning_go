package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	// fmt.Println(m)
	// fmt.Printf("%v\n", m)
	fmt.Printf("%T\n", m)
	
	m["Bell Labs"] = Vertex {
		40.68433, -74.39967, 
	}

	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Println(m["Bell Labs"])
}
