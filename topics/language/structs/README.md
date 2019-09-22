# Structs and struct types

> A struct is a sequence of named elements, called fields, each of which has a
> name and a type. -- [ref/spec](https://golang.org/ref/spec#Struct_types)

The struct serves as the main compound type in Go.

## Empty struct

```go
struct{}
```

## A struct with six fields

```go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
    x, y int
    u float32
    _ float32  // padding
    A *[]int
    F func()
}
```

## Struct types

We can define a struct type by using the type keyword.

```go
type Response struct {
    Status     string // e.g. "200 OK"
    StatusCode int    // e.g. 200
    Proto      string // e.g. "HTTP/1.0"
    ProtoMajor int    // e.g. 1
    ProtoMinor int    // e.g. 0
    ...
```

## Embedding

> A field declared with a type but no explicit field name is called an embedded
> field. An embedded field must be specified as a type name `T` or as a pointer
> to a non-interface type name `*T`, and `T` itself may not be a pointer type.

