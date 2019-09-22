# Constants

> Constants are declared like variables, but with the `const` keyword.
> Constants can be character, string, boolean, or numeric values. Constants
> cannot be declared using the := syntax.

Constants are simple.

> But in Go, a constant is just a simple, unchanging value, and from here on
> we're talking only about Go. (https://blog.golang.org/constants)



## Notes

* Constants are not variables.
* They exist only at compilation.
* Untyped constants can be implicitly converted where typed constants and variables can't.
* Think of untyped constants as having a Kind, not a Type.
* Explicit and implicit conversions.


```go
package main

import "fmt"

const Pi = 3.14

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
```

> Numeric constants are high-precision values. An untyped constant takes the
> type needed by its context.

```go
package main

import "fmt"

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
```

* [https://play.golang.org/p/Pwd-3y35FCi](https://play.golang.org/p/Pwd-3y35FCi)

Value boundaries apply; overflows prevent compilation.

```
constant 1267650600228229401496703205376 overflows int
```

A type can be explicitly set for a constant.

* [https://play.golang.org/p/0ElUyCZKsdx](https://play.golang.org/p/0ElUyCZKsdx)

## More on constants

* [ref/spec on constants](https://golang.org/ref/spec#Constants)
* [ref/spec on constant expressions](https://golang.org/ref/spec#Constant_expressions)
* [https://blog.golang.org/constants](https://blog.golang.org/constants)
* [dotGo 2019 - Dave Cheney - Constant Time](https://www.youtube.com/watch?v=pN_lm6QqHcw)
