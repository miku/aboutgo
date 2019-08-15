package main

import "fmt"

func main() {
	// delta := 'δ'
	var delta rune = 'δ'
	fmt.Printf("type of %c is %T (%d)\n", delta, delta, delta) // int32

	for _, c := range []rune{'ä', delta, 'ß', '\uFFFD'} {
		fmt.Printf("type of %c is %T (%d, %x)\n", c, c, c, c) // int32
	}
}
