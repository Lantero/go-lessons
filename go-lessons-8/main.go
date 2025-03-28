// Defer statements
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // LIFO (Last in, first out)
	}

	fmt.Println("done")
}
