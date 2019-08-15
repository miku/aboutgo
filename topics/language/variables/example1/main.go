package main

import "fmt"

var (
	a = 1
	b = 2
	c = "Hello"
)

func main() {
	d := 4

	var (
		e, f int
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
