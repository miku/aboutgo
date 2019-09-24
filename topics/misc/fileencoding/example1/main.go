package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Printf("%d\n", len(s))                    // 9
	fmt.Printf("%d\n", utf8.RuneCountInString(s)) // 13
}
