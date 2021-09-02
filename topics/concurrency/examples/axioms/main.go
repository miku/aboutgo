// fatal error: all goroutines are asleep - deadlock!
// panic: close of closed channel
package main

import (
	"fmt"
)

func main() {
	// s()
	// r()
	// c()
	// z()
}

// send to nil panics
func s() {
	var c chan string
	c <- "deadlock"
}

// receive from nil panics
func r() {
	var c chan string
	fmt.Println(<-c)
}

// close more than once panics
func c() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
			<-ch
			close(ch)
		}
	}()
	ch <- 1
	ch <- 2
}

// receive on closed channel return zero value of type
func z() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", <-c) // 1 2 3 0 0
	}
	fmt.Println()
}
