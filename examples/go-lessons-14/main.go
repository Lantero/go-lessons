// Interfaces
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	if v == nil {
		return 0
	}
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	var n Vertex
	var b Abser
	b = &n // An interface with nil value can invoke methods, and the method would receive nil

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v -> this would not compile

	describe(a)
	fmt.Println(a.Abs())
	describe(b)
	fmt.Println(b.Abs())

	// The empty interface is a way to define any type in go, since all types implement the empty interface
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i = "hello"
	describe(i)

	// The following lines will throw runtime error because a nil interface cannot call methods
	var c Abser
	fmt.Println(c.Abs())
}