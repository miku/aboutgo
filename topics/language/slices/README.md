# Slices

A slice is the common container type in Go - the type `[]T` is a slice with
elements of type `T`.

> A slice is a dynamically-sized, flexible view into the elements of an array.
> In practice, slices are much more common than arrays.

## Getting a slice

### From an Array

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
