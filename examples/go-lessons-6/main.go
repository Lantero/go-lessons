// For loops
package main

import "fmt"

func main() {
	// 1. Full for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 2. Init and post statements are optional
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	// 3. Now remove semicolons and get a while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// 4. Infinite loop
	// for {
	// }
}

