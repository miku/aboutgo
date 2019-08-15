package main

import "fmt"

func main() {

	for _, c := range "Hällo Wörld" {
		fmt.Printf("%c % 4d, % 4X\n", c, c, c)
	}
}
