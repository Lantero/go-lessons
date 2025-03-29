// Structs
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v.X)
	// Pointer access
	p := &v
	p.Y = 1e9 // Equivalent notation of (*p).Y = 1e9
	fmt.Println(v)
	// Ways of declaring structs
	v1 := Vertex{1, 2}  // has type Vertex
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p1 := &Vertex{1, 2} // has type *Vertex
	fmt.Println(p1, v1, v2, v3)
	// Why using struct pointers is better than using structs directly?
	// 1. Memory efficiency when passing structs to functions
	// 2. You can modify the struct when being passed to a function, otherwise it creates a copy
	// 3. Pointers can be nil, a great way to indicating absence of value
	// 4. Your struct methods are either value receivers (They do not modify original struct), or
	//    pointer receivers, which allow you to modify the struct's fields
}
