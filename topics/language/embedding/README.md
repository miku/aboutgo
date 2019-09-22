# Type embedding

Go does not strive for type hierarchies, it prefers composition and interfaces
for method dispatch.

> Code reuse is not provided by type hierarchy but via composition. Language
> ecosystems with classical inheritance is often suffering from excessive level
> of indirection and premature abstractions based on inheritance which later
> makes the code complicated and unmaintainable.

Embedding types provide the final piece of sharing and reusing state and
behavior between types. Through the use of inner type promotion, an inner type's
fields and methods can be directly accessed by references of the outer type.

## Notes

* Embedding types allow us to share state or behavior between types.
* The inner type never loses its identity.
* This is not inheritance.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* The outer type can override the inner type's behavior.

## Links

[Methods, Interfaces and Embedded Types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy
[Embedding is not inheritance](https://rakyll.org/typesystem/) - JBD

## Code Review

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/mT4iWg10YEp))
[Embedding types](example2/example2.go) ([Go Playground](https://play.golang.org/p/avo8I21N-qq))
[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/pdwB9dxD1MR))
[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/soB4QujV4Sj))

## Exercises

### Exercise 1

Copy the code from the template. Add a new type CachingFeed which embeds Feed and overrides the Fetch method.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/kdHgALCIPIs)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/FbvPJoQc4In))

----

## Combining interfaces

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
defer f.Unlock()
```

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
