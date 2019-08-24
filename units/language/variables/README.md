# Variables and Types

> A variable is a storage location for holding a value. The set of permissible
> values is determined by the variable's type.

Every value in Go has

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

Here,

### Shorthand declaration

Omitting *var* and type. Can only be used within functions.

```go
x := 123
```


