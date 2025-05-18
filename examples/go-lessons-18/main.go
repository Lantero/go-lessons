// Go routines 101
package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Simple go routine
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

// 2. Channels to synchronize go routines (Blocking)
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// 4. Range and close channels
func fibonacci_1(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 5. Select statement
func fibonacci_2(c, quit chan int) {
	x, y := 0, 1
	for {
		select { // Whichever channel is ready first wins, if both are ready, it is non-deterministic
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
		// default: // receiving from c would block, you can use default to avoid blocking
	}
}

// 6. Mutexes
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	// 1. Simple go routine
	go say("world")
	say("hello")
	fmt.Println()

	// 2. Channels to synchronize go routines (Blocking)
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	// 3. Buffered channels (Non-blocking)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // This will break the program as an overbuffered channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 4. Range and close channels
	c = make(chan int, 10)
	// c, ok := <-ch // This will pop a value and check if the channel is closed (ok = false)
	go fibonacci_1(cap(c), c)
	for i := range c { // Iterates over the channel until it is closed
		fmt.Println(i)
	}

	// 5. Select statement
	c = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_2(c, quit)

	// 6. Mutexes
	z := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go z.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(z.Value("somekey"))
}
