// Common Go interfaces
package main

import (
	"fmt"
	"image"
	"io"
	"strings"
	"time"
)

// Defined by fmt package
// type Stringer interface {
//     String() string
// }

// Built-in
// type error interface {
//     Error() string
// }

// Defined by io package
// type Reader interface {
//     Read(b []byte) (n int, err error)
// }

// Defined by image package
// type Image interface {
//     ColorModel() color.Model
//     Bounds() image.Rectangle
//     At(x, y int) color.Color
// }

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// Stringer interface
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	// error interface
	if err := run(); err != nil {
		fmt.Println(err)
	}
	// Reader interface
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
	// Image interface
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
