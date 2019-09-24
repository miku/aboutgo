# Go file encoding

All Go source files are assumed Unicode, encoded in UTF-8. Which means that the following is a valid program:

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, 世界")
}
```

> Source code in Go is defined to be UTF-8 text; no other representation is
> allowed. ([https://blog.golang.org/strings](https://blog.golang.org/strings))

## The len function

There is a built-in function [`len`](https://golang.org/pkg/builtin/#len), which works with various data types.

## What is a rune?

A `rune` is a character.

> Go 1 introduces a new basic type, rune, to represent individual Unicode code points. It is an alias for int32.

> Character literals such as 'a', '語', and '\u0345' now have default type rune,
analogous to 1.0 having default type float64. A variable initialized to a
character constant will therefore have type rune unless otherwise specified.

## How to get the type of a value?

One way is to use string formatting.

* [pkg/fmt](https://golang.org/pkg/fmt/)

## Code Review

* [String length](example1/main.go)
* [String formatting](example2/main.go)
* [String iteration](example3/main.go)
