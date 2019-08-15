package main

// Example showing a that pointers can be keys, too - which might cause confusion.

import "fmt"

func main() {
	m := make(map[*string]string)

	k1, k2 := "k1", "k2"
	m[&k1] = "Hello"
	m[&k2] = "World"

	for k, v := range m {
		fmt.Println(k, v)
	}

	k1 = "changed"

	fmt.Println(m[&k1])
}
