# Arrays

The type `[n]T` is an array of n values of type T. An array's length is part of
its type, so arrays cannot be resized.

```go
package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    // Hello World
    // [Hello World]
    // [2 3 5 7 11 13]
}
```

## Literal initialization

Explicitly setting the size.

```go
b := [2]string{"Penn", "Teller"}
```

We can also let the compiler infer the size from the literal.

```go
b := [...]string{"Penn", "Teller"}
```

## Arrays and slices


Instead of arrays, most programs will want arrays, that can grow in size. This is provided by slices.