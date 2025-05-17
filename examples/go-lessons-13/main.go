// Methods
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
type MyFloat float64

// 1. v Vertex is a receiver, a special argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 2. You can also implement this as a regular function, but this is not idiomatic
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 3. This is a pointer receiver, allowing you to modify the value of the receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 4. Non-idiomatic way to implement the Scale method
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 5. You can define methods with non-struct types as well
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// WARNING: You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int)

func main() {
	// Struct type
	v := Vertex{3, 4}
	// Struct type - Value receiver methods
	fmt.Println(v.Abs())
	fmt.Println((&v).Abs())  // Go is smart enough to call the method with both value and pointer receivers
	fmt.Println(Abs(v))  // ! However, fmt.Println(Abs(&v)) would not work with regular functions
	// Struct type - Pointer receiver methods
	v.Scale(10)  // ! Pointer receivers can also be called with values, magically translating to pointers
	(&v).Scale(10) // ! This is a more explicit way to do the same thing as before
	fmt.Println(v.Abs())
	Scale(&v, 10)  // ! Here however, using Scale(v, 10) will give us a compile error
	fmt.Println(v.Abs())
	// Non-struct type
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// PRO TIP: In general, all methods on a given type should have either value or pointer receivers,
	// but not a mixture of both. 
}
