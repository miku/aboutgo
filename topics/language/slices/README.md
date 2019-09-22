# Slices

Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

## Notes

* Slices are like dynamic arrays with special and built-in functionality.
* There is a difference between a slices length and capacity and they each service a purpose.
* Slices allow for multiple "views" of the same underlying array.
* Slices can grow through the use of the built-in function append.

## Links

* [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) - Andrew Gerrand
* [Strings, bytes, runes and characters in Go](https://blog.golang.org/strings) - Rob Pike
* [Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices) - Rob Pike
* [Understanding Slices in Go Programming](https://www.ardanlabs.com/blog/2013/08/understanding-slices-in-go-programming.html) - William Kennedy
* [Collections Of Unknown Length in Go](https://www.ardanlabs.com/blog/2013/08/collections-of-unknown-length-in-go.html) - William Kennedy
* [Iterating Over Slices In Go](https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html) - William Kennedy
* [Slices of Slices of Slices in Go](https://www.ardanlabs.com/blog/2013/09/slices-of-slices-of-slices-in-go.html) - William Kennedy
* [Three-Index Slices in Go 1.2](https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html) - William Kennedy
* [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)

## Code Review

* [Declare and Length](example1/example1.go) ([Go Playground](https://play.golang.org/p/ydOJ1GHgR_Y))
* [Reference Types](example2/example2.go) ([Go Playground](https://play.golang.org/p/WqDnss06_9E))
* [Appending slices](example4/example4.go) ([Go Playground](https://play.golang.org/p/E-NTGM6daAA))
* [Taking slices of slices](example3/example3.go) ([Go Playground](https://play.golang.org/p/rUP9grCot8J))
* [Slices and References](example5/example5.go) ([Go Playground](https://play.golang.org/p/D88zzGYanvX))
* [Strings and slices](example6/example6.go) ([Go Playground](https://play.golang.org/p/1RntHk6UPA5))
* [Variadic functions](example7/example7.go) ([Go Playground](https://play.golang.org/p/rUjWVBMmxgP))
* [Range mechanics](example8/example8.go) ([Go Playground](https://play.golang.org/p/d1wToBg6oUu))

## Advanced Code Review

* [Three index slicing](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/2CM_LPBnfIR))

## Exercises

### Exercise 1

**Part A** Declare a nil slice of integers. Create a loop that appends 10 values to the slice. Iterate over the slice and display each value.

**Part B** Declare a slice of five strings and initialize the slice with string literal values. Display all the elements. Take a slice of index one and two and display the index position and value of each element in the new slice.

* [Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/7GfB3NOwu_c)) |
* [Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/0xv7GTHHIR_K))

## Additional notes on slices

### Create a slice from an Array

With *slice notation*, specifying a low and high bound:

```go
a[low : high]
```

* [moretypes/7](https://tour.golang.org/moretypes/7)

```go
package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)
}
```

### Declaration

The underlying (backing) array is allocated implicitly.

```go
var (
    hits   []int
    temps  []float64
    ss     []string
    matrix [][]float64
)
```

### Literal initialization

Initialization with values.

```go
var s = []string{"Berlin", "Munich", "Hamburg", "Cologne"}
```

### Zero value

The zero value for a slice is `nil`.

### Length and capacity

A slice has two properties:

* length (number of elements)
* cap (sum of length and the length of the array beyond the slice)

The builtin functions `len` and `cap` can be used to inspect the values.  The len and cap functions will both return 0 for a nil slice.

### Initialization with make

If we want to control length and capacity of the slice, we can use the builtin make function.

```go
package main

import "fmt"

var (
    s = make([]string, 10)
    t = make([]string, 10, 100)
    u []string
)

func sliceInfo(s []string) string {
    return fmt.Sprintf("len=%d, cap=%d", len(s), cap(s))
}

func main() {
    fmt.Println(sliceInfo(s))
    fmt.Println(sliceInfo(t))
    fmt.Println(sliceInfo(u))
}

// len=10, cap=10
// len=10, cap=100
// len=0, cap=0
```

## Slicing

We can create a slice from an array. But we can also create a slice from
another slice. We need to keep in mind, that the underlying array will be the
same.

```go
package main

import "fmt"

func main() {
    a := []int{0, 1, 2, 3, 4, 5, 6, 7}
    b := a[2:6]
    c := a[4:]

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    c[0] = 1000

    fmt.Println(a)

    // [0 1 2 3 4 5 6 7]
    // [2 3 4 5]
    // [4 5 6 7]
    // [0 1 2 3 1000 5 6 7]
}
```

The start and end indices of a slice expression are optional; they default to
zero and the slice's length respectively. Hence you can write the following:

```go
    ...
    a := []int{0, 1, 2, 3, 4, 5, 6, 7}
    b := a[2:6]
    c := a[4:]
    d := a[:]
    ...
```

## Growing a slice

Slices are dynamically sized, so how do we extend them? With the builtin `append` function.

```go
var s []string
s = append(s, "a")
s = append(s, "b")
```

To understand `append`, we need to recap what a slice is: A slice points to an
underlying array, which is fixed in size. What if our append operation would
require a larger underlying array?

We would need to allocate a new array, copy over all existing elements and then
insert the new one. This is exactly what `append` does.

## Concatenating slices

The `append` function allows to concatenate two slices as well.

```go
package main

import "fmt"

func main() {

    a := []string{"a", "b", "c"}
    b := []string{"d", "e", "f"}

    c := append(a, b...)
    fmt.Println(c)
    // [a b c d e f]
}
```

## Slice internals

Slices can be confusing at first.

* [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals)

A slices is a lightweight descriptor (implemented as struct), that contains:

* a pointer to an array
* length of a segment
* capacity, which is maximum length of the segment

![](https://blog.golang.org/go-slices-usage-and-internals_slice-struct.png)

A variable created with

```go
make([]byte, 5)
```

would look like this:

![](https://blog.golang.org/go-slices-usage-and-internals_slice-1.png)


## Examples

```
$ go run slices3.go
Expression                          |Value              |Nil?   |Type   |Length  |Capacity
a := []int{0, 1, 2, 3, 4, 5, 6, 7}  |[0 1 2 3 4 5 6 7]  |false  |[]int  |8       |8
b := a[2:6]                         |[2 3 4 5]          |false  |[]int  |4       |6
c := a[4:]                          |[4 5 6 7]          |false  |[]int  |4       |4
d := a[:]                           |[0 1 2 3 4 5 6 7]  |false  |[]int  |8       |8
e := c[1:]                          |[5 6 7]            |false  |[]int  |3       |3
```
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
