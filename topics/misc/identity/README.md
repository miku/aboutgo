# Truthiness and identity

## Truth

There is a type `bool` with two predeclared identiers: `true` and `false`. The
logical operators are `!`, `||` and `&&`.

## Type Identity

> A defined type is always different from any other type. Otherwise, two types
> are identical if their underlying type literals are structurally equivalent;
> that is, they have the same literal structure and corresponding components
> have identical types.

> https://golang.org/ref/spec#Type_identity

## Struct Identity

Can structs be compared?

```go
package main

import (
	"fmt"
)

type Thing struct {
	Value int
}

type Complex struct {
	S string
	M map[string]string
}

func main() {
	a, b := Thing{1}, Thing{1}
	fmt.Println(a == b)   // true
	fmt.Println(&a == &b) // false
	fmt.Println(&a == &a) // true

	a, b = Thing{1}, Thing{2}
	fmt.Println(a == b) // false

	c, d := Complex{}, Complex{}
	fmt.Println(c == d) // invalid operation: c == d (struct containing map[string]string cannot be compared)
}
```

* [https://play.golang.org/p/bX9rzmF5Tdz](https://play.golang.org/p/bX9rzmF5Tdz)
* [How to compare struct, slice, map are equal?](https://stackoverflow.com/questions/24534072/how-to-compare-struct-slice-map-are-equal)