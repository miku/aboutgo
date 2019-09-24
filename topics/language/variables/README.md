# Variables and Types

> A variable is a storage location for holding a value. The set of permissible
> values is determined by the variable's type.

Go is mostly statically typed (interfaces can accomodate multiple types).

## Variable declaration

There are four variants to introduce a variable into a program.

```go
var x int
var x int8 = 127
var x = 127
x := 127
```

### Declaration

The `var` statement declares a *list of variables*; as in function argument lists, the type is last.

```go
var x int
```

Declare multiple variables of a type at once:

```go
var x, y, z int
```

The var statement can be used outside functions. Variables declared without an
explicit initial value are given their [zero value](https://golang.org/ref/spec#The_zero_value).

> Each element of such a variable or value is set to the zero value for its
> type: false for booleans, 0 for numeric types, "" for strings, and nil for
> pointers, functions, interfaces, slices, channels, and maps. This
> initialization is done recursively.

### Declaration and initialization

Declare and initialize.

```go
var x int = 127
var x, y, z int = 1, 2, 3
```

### Declaration with type inference

The type can be omitted in a *var* statement.

```go
var x = 123
```

The type of `x` will be an `int` in this case.

### Shorthand declaration

Omitting *var* and type. Can only be used within functions.

```go
x := 123
```

## Types

> A type determines a set of values together with operations and methods
> specific to those values. A type may be denoted by a type name, if it has
> one, or specified using a type literal, which composes a type from existing
> types.

Note: An example for type without name would be an anonymous struct.

Type can be roughly put into the following [groups](https://golang.org/ref/spec#Types):

* Boolean types &mdash; bool
* Numeric types &mdash; byte, int, uint, int8, int16, ..., float32, float64
* String types &mdash; string
* Array types &mdash; [n]T
* Slices types &mdash; []T
* Struct types &mdash; T struct
* Pointer types &mdash; *T
* Function types &mdash; func F()
* Interface types &mdash; method set
* Map types &mdash; map[K]V
* Channel types &mdash; chan(T)

> The language predeclares certain type names.

For example, `string`, `int32`, `rune` are predeclared types.

> Others are introduced with [type declarations](https://golang.org/ref/spec#Type_definitions).
> A type definition creates a new, distinct type with the same *underlying
> type* and operations as the given type, and binds an identifier to it.

```go
package main

import "fmt"

func main() {
    type Temperature float64
    var x, y Temperature

    x = 1.0
    y = 2.0
    fmt.Println(x + y) // 3
}
```

Note that there are not classes, class definitions or objects.

> … the more important idea is the separation of concept: data and behavior are
> two distinct concepts in Go, not conflated into a single notion of “class”.
> -- [Rob Pike](https://twitter.com/rob_pike/status/942528032887029760)

Emphasis is on data, e.g. using structs. Inheritance is possible, but a less
central concept.

## Zero Value

There are not uninitialized variables in Go, every type has a zero value. You
cannot distringuish between a variable declared `int` and a variable that is 0.
This can occassionally lead to problems.

One example would be a workaround, where you work with pointers (where the zero
value is nil) - but that leads to other obstacles, as this [issue in the AWS
SDK](https://github.com/aws/aws-sdk-go/issues/114) illustrates.

## Code Review

* [Variables](example1/main.go)