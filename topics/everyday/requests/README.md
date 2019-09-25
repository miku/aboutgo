# Making requests

The [net/http](https://golang.org/pkg/net/http/) package contains HTTP server and client code.

## Making a GET request

The package comes with a default client, that is available directly - it will
return a `http.Response` and an error.

```
resp, err := http.Get("http://example.com/")
```

## Code Review

* [Example 1](example1/main.go)
* [Example 2](example2/main.go)
* [Example 3](example3/main.go)
* [Example 4](example4/main.go)
* [Example 5](example5/main.go)

## Exercises

