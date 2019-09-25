# Services

## Handler

There are a couple of concepts related to services.

* [http.Handler](https://golang.org/pkg/net/http/#Handler)

The core interface for handline HTTP request.

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

## Handlerfunc

You can turn a function (with the right signature) into a type, that satisfies
http.Handler.

* [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc)

> The HandlerFunc type is an adapter to allow the use of ordinary functions as
> HTTP handlers. If f is a function with the appropriate signature,
> HandlerFunc(f) is a Handler that calls f.

With first class functions, it is possible to start an HTTP handler with a
single (long) line.

## HandleFunc

If we type `http.HandleFunc`, we are working on the `DefaultServeMux`, a
[http.ServeMux](https://golang.org/pkg/net/http/#ServeMux) added for convenience
in the http package.

> ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
> request against a list of registered patterns and calls the handler for the
> pattern that most closely matches the URL.

The `ServeMux` is itself an HTTP handler, as it implements
[ServeHTTP](https://golang.org/pkg/net/http/#ServeMux.ServeHTTP).




## Exercises

