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

