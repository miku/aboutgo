# Code and Naming Conventions

## Always format your code

Always go fmt your code. Or use goimports (superset of go fmt).

## Do not stutter

If you have a package `time` and a function `TimeSince` you should keep it brief by naming your function just `Since` so the user of your code can write:

```go
time.Since(...)
```

* [Package names](https://blog.golang.org/package-names)

## Short variable names

It is not unreasonable to follow conventions and to name an `io.Writer` just
`w`. An excerpt for `src/io/io.go`.

```go
// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
func WriteString(w Writer, s string) (n int, err error) {
	if sw, ok := w.(StringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}
```

## Package names

Try to use simple names.

> A package names with under_scores, hy-phens or mixedCaps should be avoided.

## Keep an eye on your code

The `go vet` tool is available out of the box.

> Vet examines Go source code and reports suspicious constructs, such as Printf
> calls whose arguments do not align with the format string.

It is also part of the Hackathon project, [Go Report
Card](https://goreportcard.com/), along with a few other checks.

