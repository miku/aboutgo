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

