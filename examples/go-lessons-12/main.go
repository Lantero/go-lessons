// Maps
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m0 map[string]Vertex
var m1 map[string]Vertex
var m2 map[string]Vertex
var m3 map[string]Vertex

func main() {
	// 1. Define maps
	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m2 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	m3 = map[string]Vertex{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m0)
	fmt.Println(m0 == nil)
	fmt.Println(m1["Bell Labs"])
	fmt.Println(m2)
	fmt.Println(m3)

	// 2. Mutate maps
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
