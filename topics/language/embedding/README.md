# Type embedding

Go does not strive for type hierarchies, it prefers composition and interfaces
for method dispatch.

> Code reuse is not provided by type hierarchy but via composition. Language
> ecosystems with classical inheritance is often suffering from excessive level
> of indirection and premature abstractions based on inheritance which later
> makes the code complicated and unmaintainable.

The closest notion to inheritance in Go is type embedding.

The io package derives various new types by simply embedding existing types.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

The `ReadWriter` interface will be fullfilled by type implementing both `Read`
and `Write` with the proper signature.

> There's an important way in which embedding differs from subclassing. When we
> embed a type, the methods of that type become methods of the outer type, but
> when they are invoked the receiver of the method is the inner type, not the
> outer one. In our example, when the Read method of a bufio.ReadWriter is
> invoked, it has exactly the same effect as the forwarding method written out
> above; the receiver is the reader field of the ReadWriter, not the ReadWriter
> itself. -- [Effective Go](https://golang.org/doc/effective_go.html#embedding)

## Example: Embedding a lock

A use case for embedding is including a lock in a struct. Instead of declaring
named field, we embed a `sync.Mutex` directly.

```go
type File struct {
    sync.Mutex
    rw io.ReadWriter
}
```

File object can then `Lock` and `Unlock` themselves directly.

```go
f := File{}
f.Lock()
```

