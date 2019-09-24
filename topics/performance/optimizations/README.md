# Optimizations

> Programmers waste enormous amounts of time thinking about, or worrying about,
> the speed of noncritical parts of their programs, and these attempts at
> efficiency actually have a strong negative impact when debugging and
> maintenance are considered. We should forget about small efficiencies, say
> about 97% of the time: premature optimization is the root of all evil. -- Donald Knuth

## Profile

Try to include profiling in your workflow.

* [wiki/Performance](https://github.com/golang/go/wiki/Performance)

Identify bottlenecks first.

## Generic Ideas

### Reduce allocations

Consider the following two functions, one might be more efficient, with respect
to allocations.

```
func (r *Reader) Read() ([]byte, error)
func (r *Reader) Read(buf []byte) (int, error)
```

### Avoid string concatenation

Go strings are immutable. The following might not be the most efficient way to build a string.

```go
s := request.ID
s += " " + client.Addr().String()
s += " " + time.Now().String()
r = s
```

Other variants:

```go
var b bytes.Buffer
fmt.Fprintf(&b, "%s %v %v", request.ID, client.Addr(), time.Now())
r = b.String()
```        

Using the [string.Builder](https://golang.org/pkg/strings/#Builder):

```go
var b strings.Builder
b.WriteString(request.ID)
b.WriteString(" ")
b.WriteString(client.Addr().String())
b.WriteString(" ")
b.WriteString(time.Now().String())
r = b.String()
```

### Parallelize

If the problem is parallelizable, it might be worth the effort to improve performance.

Example for a generic [parallel line processor](https://github.com/miku/parallel).

When using goroutines, consider batching data - otherwise the communication
overhead might exceed the benefits.
