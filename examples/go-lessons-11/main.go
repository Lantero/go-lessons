// Arrays and slices
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Example using arrays, which cannot be resized (They are static)
	fmt.Println("// 1. Arrays (static)")
	var hello [2]string
	hello[0] = "Hello"
	hello[1] = "World"
	fmt.Println(hello[0], hello[1])
	fmt.Println(hello)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Printf("primes %T\n\n", primes)

	// 2. Slice as a dynamically size view of the array
	fmt.Println("// 2. Slices (dynamic)")
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Printf("s: %T\n\n", s)

	// 3. The underlying data is the same
	fmt.Println("// 3. Modifying the same underlying data with slices")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	fmt.Println()

	// 4. Slice literals allow you to create an array and the slice view on top, all together,
	//    without explicitely assigning the size in the declaration
	fmt.Println("// 4. Slice literals")
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(st)
	fmt.Println()

	// 5. You can omit high or low bounds when defining slices
	fmt.Println("// 5. Slice defaults")
	sh := []int{2, 3, 5, 7, 11, 13}
	sh = sh[1:4]
	fmt.Println(sh)
	sh = sh[:2]
	fmt.Println(sh)
	sh = sh[1:]
	fmt.Println(sh)
	fmt.Println()

	// 6. Slice length and capacity
	//    len() gives you the size of the slice
	//    cap() gives you the size of the underlying array, starting with the first slice element
	fmt.Println("// 6. Slice length and capacity")
	printSlice := func(s []int) {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
	slc := []int{2, 3, 5, 7, 11, 13}
	printSlice(slc)
	// Slice the slice to give it zero length.
	slc = slc[:0]
	printSlice(slc)
	// Extend its length.
	slc = slc[:4]
	printSlice(slc)
	// Drop its first two values.
	slc = slc[2:]
	printSlice(slc)
	fmt.Println()

	// 7. Zero value of a slice is nil, and has 0 length and capacity
	fmt.Println("// 7. Slice zero value")
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
	fmt.Println()

	// 8. Initializing a slice with make
	fmt.Println("// 8. Create a slice and initialize the contents with zero values")
	ma := make([]int, 5) // Second argument is length
	printSlice(ma)
	mb := make([]int, 0, 5) // Third optional argument is capacity
	printSlice(mb)
	mc := mb[:2]
	printSlice(mc)
	md := mc[2:5]
	printSlice(md)
	fmt.Println()

	// 9. You can create slices of slices
	fmt.Println("// 9. Slices of slices")
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println()

	// 10. Appending to an slice
	fmt.Println("// 10. Appending to slices")
	var sa []int
	printSlice(sa)
	// append works on nil slices, a new array is allocated if the existing one does not fit
	sa = append(sa, 0)
	printSlice(sa)
	// The slice grows as needed.
	sa = append(sa, 1)
	printSlice(sa)
	// We can add more than one element at a time.
	sa = append(sa, 2, 3, 4)
	printSlice(sa)
	fmt.Println()

	// 11. Range
	fmt.Println("// 11. Use range to iterate over slices")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
