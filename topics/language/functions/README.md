# Functions

Functions are at the core of the language. They provide a mechanism to group and
organize our code to separate and distinct pieces of functionality. They can be
used to provide an API to the packages we write and are a core component to
concurrency.

## Notes

* Functions can return multiple values and most return an error value.
* The error value should always be checked as part of the programming logic.
* The blank identifier can be used to ignore return values.

Functions are declared with the `func` keyword.

```go
func Inc(a int) int {
    return a + 1
}
```

It is possible to group parameters of the same type.

```go
func Sum(a int, b int) int {
    return a + b
}

func Sum(a, b int) int {
    return a + b
}
```

## Signature

There is no function overloading in Go, a function has a name and this name can
only be used for a single implementation. This means, that you cannot write the
following.

```go
func Inc(a int) int {
    return a + 1
}

func Inc(a float64) float64 {
    return a + 1
}
```

This would yield an `Inc redeclared in this block` error at compile time.

> Functional Names & Immutable Meanings

> Twenty years ago, Rob Pike and I were modifying the internals of a Plan 9 C
> library, and Rob taught me the rule of thumb that when you change a
> function’s behavior, you also change its name. The old name had one meaning.
> By using a different meaning for the new name and eliminating the old one, we
> ensured the compiler would complain loudly about every piece of code that
> needed to be examined and updated, instead of silently compiling incorrect
> code. -- [Functional Names & Immutable
> Meanings](https://research.swtch.com/vgo-import#functional_names_immutable_meanings)

## Multiple return values

> One of Go's unusual features is that functions and methods can return
> multiple values.

This is a take on the in-band error signals in C and the problem it creates.

> In C, a write error is signaled by a negative count with the error code
> secreted away in a volatile location. In Go, **Write can return a count and an
> error**.

```go
func (file *File) Write(b []byte) (n int, err error)
```

## Named result parameters

> The return or result "parameters" of a Go function can be given names and
> used as regular variables, just like the incoming parameters. When named,
> they are initialized to the zero values for their types when the function
> begins; if the function executes a return statement with no arguments, the
> current values of the result parameters are used as the returned values.

Example.

```go
func nextInt(b []byte, pos int) (value, nextPos int) {
```

This can be both useful and confusing. It should be used with care and when in
doubt, omitted.


## Links

* [https://golang.org/doc/effective_go.html#functions](https://golang.org/doc/effective_go.html#functions)

## Code Review

* [Return multiple values](example1/example1.go) ([Go Playground](https://play.golang.org/p/-7A-lGLv2TK))
* [Blank identifier](example2/example2.go) ([Go Playground](https://play.golang.org/p/ID54tVxM5B0))
* [Redeclarations](example3/example3.go) ([Go Playground](https://play.golang.org/p/EDRhDh2r1Mj))
* [Anonymous Functions/Closures](example4/example4.go) ([Go Playground](https://play.golang.org/p/h8Yi_2Sxsmu))

## Advanced Code Review

* [Recover panics](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/Wcd3CLbQZBH))

## Exercises

### Exercise 1

**Part A** Declare a struct type to maintain information about a user. Declare a function that creates value of and returns pointers of this type and an error value. Call this function from main and display the value.

**Part B** Make a second call to your function but this time ignore the value and just test the error value.

* [Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/5vEQxEzq3i_D)) |
* [Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/P8wC324WWuh))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
