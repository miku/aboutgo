# First Go program

## Verify your Go installation

```
$ go version
go version go1.12.7 linux/amd64
```

The output should show [version](https://golang.org/doc/devel/release.html), [platform](https://github.com/golang/go/wiki/MinimumRequirements#operating-systems) and [architecture](https://github.com/golang/go/wiki/MinimumRequirements#architectures).

```
$ go version
go version go1.12.6 linux/arm
```

## Your main.go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

## Run

You can *just* run your code with `go run`, e.g.

```
$ go run main.go
Hello World
```

Behind the scenes, the go tool compiles the file to a temporary location, then executes it. You can pass flags to run to see what is happening.

```
$ go run -work main.go 
WORK=/tmp/go-build907419825
Hello World
```

Even more information is show with the `-x` flag.

## Build

To build an executable, use the `build` subcommand:

```
$ go build main.go
```

By default, the binary will be named as the program, e.g. `main` or `main.exe`
on Windows.

This can be adjusted with the `-o` flag:

```
$ go build -o myprog main.go
```

The result is a statically linked binary.

```
$ file main
main: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked

$ ldd main
        not a dynamic executable
```

> Static linking is the result of the linker copying all library routines used
> in the program into the executable image. This may require more disk space and
> memory than dynamic linking, but is both faster and more **portable**, since it
> does not require the presence of the library on the system where it is run.

There are cases, where the presence of libraries is required.
