# Package io

Package [io](https://golang.org/pkg/io/) defines helpers for working with data
and data streams.

Especially [io.Reader](https://golang.org/pkg/io/#Reader) and
[io.Writer](https://golang.org/pkg/io/#Writer) are widely used interfaces.


## Flexible API

What could be improved in the following function?

```go
func Contains(f *os.File, term string) (bool, error) {
    b, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, err
    }
    return bytes.Contains(b, term), nil
}
```