# Working with files

## Opening files

```go
f, err := os.Open("README")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

Files implement a couple of interfaces, e.g. `io.Reader`, `io.Writer` and a few more.


