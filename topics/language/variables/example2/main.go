package main

import "fmt"

type celsius float64

var (
	freezingC celsius = 0
	boilingC  celsius = 100
)

func main() {
	var current = 26.0
	fmt.Printf("%v %T\n", freezingC, freezingC)

	// diff := BoilingC - current // invalid operation: BoilingC - current (mismatched types Celsius and float64)
	diff := boilingC - celsius(current)
	fmt.Printf("%v %T\n", diff, diff)
}
