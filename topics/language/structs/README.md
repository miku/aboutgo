## Struct Types

Struct types are a way of creating complex types that group fields of data together. They are a great way of organizing and sharing the different aspects of the data your program consumes.

A computer architecture’s potential performance is determined predominantly by its word length (the number of bits that can be processed per access) and, more importantly, memory size, or the number of words that it can access.

## Notes

* We can use the struct literal form to initialize a value from a struct type.
* The dot (.) operator allows us to access individual field values.
* We can create anonymous structs.

## Quotes

_"Implicit conversion of types is the Halloween special of coding. Whoever thought of them deserves their own special hell." - Martin Thompson_

## Links

* [Understanding Type in Go](https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html) - William Kennedy
* [Object Oriented Programming in Go](https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html) - William Kennedy
* [Padding is hard](https://dave.cheney.net/2015/10/09/padding-is-hard) - Dave Cheney
* [Structure Member Alignment, Padding and Data Packing](https://www.geeksforgeeks.org/structure-member-alignment-padding-and-data-packing/)
* [The Lost Art of Structure Packing](http://www.catb.org/esr/structure-packing) - Eric S. Raymond

## Code Review

* [Declare, create and initialize struct types](example1/example1.go) ([Go Playground](https://play.golang.org/p/djzGT1JtSwy))
* [Anonymous struct types](example2/example2.go) ([Go Playground](https://play.golang.org/p/09cxjnmfcdC))
* [Named vs Unnamed types](example3/example3.go) ([Go Playground](https://play.golang.org/p/ky91roJDjir))

## Advanced Code Review

* [Struct type alignments](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/rAvtS7cgD0z))

## Exercises

### Exercise 1

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

* [Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/h-7BEn2U3Rz)) |
* [Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/eT_gLZKeHk-))


# Additional notes

> A struct is a sequence of named elements, called fields, each of which has a
> name and a type. -- [ref/spec](https://golang.org/ref/spec#Struct_types)

The struct is the main construct to build custom types.

A struct contains zero or more fields. Access to fields is through dot notation.

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

Accessing fields:

```go
var resp Response
fmt.Println(resp.Status)
```

## Anonymous structs

Example of an anonymous struct to pass data into a html template.

```go
// title and users have been declared and populated before.
data := struct {
    Title string
    Users []*User
}{
    title,
    users,
}
err := tmpl.Execute(w, data)
```

They are also useful, when writing [table-driven tests](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go).

## Embedding

> A field declared with a type but no explicit field name is called an embedded
> field. An embedded field must be specified as a type name `T` or as a pointer
> to a non-interface type name `*T`, and `T` itself may not be a pointer type.

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
