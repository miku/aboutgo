package main

import "fmt"

// Exercise: Declare a few variables of different types. Use different declarations. Try to add two different numeric types.

var a = 1
var b = 1.0

func main() {
	var c = 1.0
	d := "Hello"

	fmt.Println(c + float64(a))
	fmt.Println(d)
}
