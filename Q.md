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
