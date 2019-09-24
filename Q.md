# Questions

## Does tooling break during the move to Go modules?

Yes. For progress, see: [#24661](https://github.com/golang/go/issues/24661).

## Should I use named return parameters?

Only in short functions.

> Naked return statements should be used only in short functions, as with the
> example shown here. They can harm readability in longer functions. -- https://tour.golang.org/basics/7

## Where is the module cache in Golang?

> In Go 1.11 and later, it default to `$GOPATH/pkg/mod`.

* [https://stackoverflow.com/questions/52126923/where-is-the-module-cache-in-golang](https://stackoverflow.com/questions/52126923/where-is-the-module-cache-in-golang)

For more information, see:

```
$ go help modules
```

## Do you use new or ampersand (&) on structs?

* [Golang basics struct and new() keyword](https://stackoverflow.com/q/34543430/89391)
* [Why is there a “new” in Go?](https://softwareengineering.stackexchange.com/q/210399/436)
* [Why is there a “new” keyword in Go?](https://groups.google.com/forum/#!topic/golang-nuts/K3Ys8qpml2Y)

> [new](https://golang.org/doc/effective_go.html#allocation_new) is a built-in
> function that *allocates* memory, but unlike its namesakes in some other
> languages it does not initialize the memory, it only zeros it.

The [address operator](https://golang.org/ref/spec#Address_operators) [...]
generates a pointer of type `*T` to `x`.

> Also, although this is a lesser point, the meaning of new in Go is not
the same as in some other languages. All it does is allocate; it does not run
any type-specific constructor. --
[https://groups.google.com/forum/#!msg/golang-nuts/kWXYU95XN04/iRfB7YEt57UJ](https://groups.google.com/forum/#!msg/golang-nuts/kWXYU95XN04/iRfB7YEt57UJ)

## Does Go distinguish between stack and heap?

Yes and no. Yes, since both are used and no, since it does not play a role when
writing Go.

> Again, the whole concept of heap is not part of the Go
language. -- [https://groups.google.com/d/msg/golang-nuts/kWXYU95XN04/vthkrDeId2cJ](https://groups.google.com/d/msg/golang-nuts/kWXYU95XN04/vthkrDeId2cJ)

## How to get the size of an value?

There is Sizeof function in the unsafe package.

* [playground](https://play.golang.org/p/pg6vJSYMqVd)

## How are types compared?

* Basic data types are always comparable using the == and != operators: integer
values, floating-point numbers, complex numbers, boolean values, string values,
constant values.

* Array values are comparable, if they contain a comparable element type
* Pointer values are comparable.
* Channel values are comparable.
* Interface values are comparable.
* Comparing interface values works only if the dynamic type is comparable.
* Function values, Slice values and Map values are not comparable, they can
  only be compared with nil, as a special case.

> Two struct values can be tested for equality by comparing the values of their
> individual fields. In general, two struct values are considered equal if they
> are of the same type and the their corresponding fields are equal.

> Two pointer values are considered equal if they point to the same value in
> memory (or if they are nil).

> Two interface values are considered equal if their underlying concrete types
> and their values are comparable and are equal or if both interfaces are nil.

> Two channel values are considered equal if they originated from the same make call.


