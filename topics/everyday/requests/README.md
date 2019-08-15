# Making requests

The [net/http](https://golang.org/pkg/net/http/) package contains HTTP server and client code.

## Making a GET request

The package comes with a default client, that is available directly - it will
return a `http.Response` and an error.

```
resp, err := http.Get("http://example.com/")
```