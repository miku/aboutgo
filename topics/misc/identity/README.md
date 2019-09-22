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

func main() {
	a, b := Thing{1}, Thing{1}
	fmt.Println(a == b)   // true
	fmt.Println(&a == &b) // false
	fmt.Println(&a == &a) // true

	a, b = Thing{1}, Thing{2}
	fmt.Println(a == b) // false

}
```

* [https://play.golang.org/p/fsZ9_O6cG-F](https://play.golang.org/p/fsZ9_O6cG-F)