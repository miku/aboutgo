# Packages

> Go programs are constructed by linking together packages. A package in turn is
> constructed from one or more source files that together declare constants,
> types, variables and functions belonging to the package and which are
> accessible in all files of the same package. Those elements may be exported
> and used in another package.

The `package` clause begins each source file and defines to which package the
file belongs.

> An implementation may require that all source files for a package inhabit the
> same directory.

The current Go implementation requires it.

## The main package

There is one special package called `main`,  which contains a `main` function,
which serves as entry point for executable programs.

By convention, many projects group their executables in to a separate directory,
called `cmd` plus one subdirectory for each executable. An example from the
[perkeep](https://github.com/perkeep/perkeep/tree/master/cmd) project:

```
perkeep/cmd
├── [4.0K]  pk
│   ├── [1.7K]  camtool.go
│   ├── [1.9K]  claims.go
|   ...
├── [4.0K]  pk-deploy
│   ├── [ 809]  camdeploy.go
│   └── [5.4K]  gce.go
├── [4.0K]  pk-devimport
│   └── [5.0K]  devimport.go
├── [4.0K]  pk-get
│   ├── [1.1K]  doc.go
│   ├── [ 12K]  get.go
│   └── [3.1K]  graph.go
├── [4.0K]  pk-mount
│   ├── [3.6K]  doc.go
│   ├── [6.5K]  pkmount.go
│   └── [ 711]  pkmount_other.go
└── [4.0K]  pk-put
    ├── [1.1K]  androidx.go
    ├── [2.5K]  attr.go
    ...
6 directories, 55 files

```

## Historical note on GOPATH

The GOPATH, by default `$HOME/go` organized third-party Go code (own and
dependencies) in a single tree. Naming was tied to directories. This worked
relatively fine, surprisingly - but had also lots of problems. Problems, that
Google admittedly solves using internal tools.

With Go 1.11 (August 2018) preliminary support for [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more).

> Go 1.11 adds preliminary support for a new concept called “modules,” an
> alternative to GOPATH with integrated support for versioning and package
> distribution. Using modules, developers are no longer confined to working
> inside GOPATH, version dependency information is explicit yet lightweight, and
> builds are more reliable and reproducible.

Currently (2019), the Go community is in a transistion process with regard to
dependency management. Many projects already using Go modules, but not all.
Rough edges remain, unfortunately.

However, the theoretical underpinnings have been documented in depth in a number
of blog posts:

* [Go & Versioning](https://research.swtch.com/vgo) (all blog posts as a [single
  PDF](https://github.com/golang-leipzig/gomodintro/blob/master/vgo-all.pdf))

However, the `$GOPATH` is [still relevant](https://github.com/golang/go/wiki/GOPATH), it is not not used for tracking dependencies anymore.

## Internal packages

Since Go 1.5, there is the mechanism to hide packages and restrict their importability - called internal packages.

> An import of a path containing the element “internal” is disallowed if the
> importing code is outside the tree rooted at the parent of the “internal”
> directory.

For example:

* Code in `/a/b/c/internal/d/e/f` can be imported only by code in the directory
  tree rooted at `/a/b/c`. It cannot be imported by code in `/a/b/g`.
* `$GOROOT/src/pkg/internal/xxx` can be imported only by other code in the
  standard library (`$GOROOT/src/`).
* `$GOROOT/src/pkg/net/http/internal` can be imported only by the standard
  `net/http` and `net/http/*` packages.
* `$GOPATH/src/mypkg/internal/foo` can be imported only by code in
  `$GOPATH/src/mypkg`.

Code in internal packages is public, documented, but cannot be imported:

```go
package main

import (
        "fmt"

        "perkeep.org/internal/azure/storage"
)

func main() {
        var c storage.Client
        fmt.Println(c)
}
```

When trying to run this program, we get an:

```
$ go run main.go
main.go:6:2: use of internal package perkeep.org/internal/azure/storage not allowed
```

One motivation for internal packages from the [design document](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU):

> There is unexported code duplicated in the standard library because sharing it
> would have required exporting those functions: net/http and net/http/httputil
> share a few parsing functions, and there are four ‘func itoa’ in various
> packages. Similarly, cmd/nm and cmd/objdump contain the same file reading
> code.
