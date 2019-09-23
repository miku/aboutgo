# Error handling

Go has not exceptions. [Errors are values](https://blog.golang.org/errors-are-values).

The typical idiom looks like:

```go
v, err := myFunc()
if err != nil {
    // Handle error, report to caller or fail.
}
```

The standard library provides an predeclared type called `error` - it's definition is:

```go
type error interface {
	Error() string
}
```

As of 2019, work is in progress to allow for finer grained errors. They are
possible today, see for example the
[os.PathError](https://golang.org/pkg/os/#PathError), but their usage can be
simplified. There is a new package [errors](https://golang.org/pkg/errors/) in
the standard library that contains functions to manipulate errors.

More examples can be found at the [errors package
examples](https://golang.org/pkg/errors/).

