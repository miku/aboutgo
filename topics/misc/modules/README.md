# Go Modules

After a few years with a variety of approaches to dependency management (e.g.
vendoring) as of 2018, Go modules start to gain adoption. It introduces a couple
of [new concepts](https://github.com/golang/go/wiki/Modules#new-concepts).

> A module is a collection of related Go packages that are versioned together as
> a single unit.

> Modules record precise dependency requirements and create reproducible builds.

> Most often, a version control repository contains exactly one module defined
in the repository root.

The relationsships between repositories, modules and packages:

Summarizing the relationship between repositories, modules, and packages:

* A repository contains one or more Go modules.
* Each module contains one or more Go packages.
* Each package consists of one or more Go source files in a single directory.

## The go.mod file

A module is defined by a tree of Go source files with a `go.mod` file in the
tree's root directory. Module source code may be located outside of `GOPATH`.
There are four directives: `module`, `require`, `replace`, `exclude`.

A basic example:

```
module github.com/my/thing

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency/v4 v4.0.0
)
```

A module declares its **identity** in its go.mod via the module directive, which
provides the module path. The import paths for all packages in a module share
the module path as a common prefix. The module path and the relative path from
the go.mod to a package's directory together determine a package's import path.

More on `go.mod` in the [wiki](https://github.com/golang/go/wiki/Modules#gomod).

> Each dependency requirement is written as a module path and a specific
> semantic version.

A comment can indicate, whether a dependency is not used by module directly.

## The go.sum file

The go.sum contains checksums to verify code equality. Nothing gets deleted from
this file, it only grows.

## Creating a go.mod file

```
$ go mod init
```

> Init initializes and writes a new go.mod to the current directory, in effect
creating a new module rooted at the current directory. The file go.mod must not
already exist. If possible, init will guess the module path from import comments
(see 'go help importpath') or from version control configuration. To override
this guess, supply the module path as an argument.

Example workflow:

* create a repository on github.com, clone, then run `go mod init`

## Integration with the go tool

The `go` tool is aware of modules. It looks at the `go.mod` file and it's local
package cache to determine, whether all dependencies are available. There is no
extra step to run, just run `go run` or `go test` as you would usually would.

The module cache is located under `$GOPATH/pkg/mod`, but that is an
implementation detail.


> When using modules, GOPATH is no longer used for resolving imports. However,
> it is still used to store downloaded source code (in GOPATH/pkg/mod) and
> compiled commands (in GOPATH/bin). (https://golang.org/cmd/go/#hdr-GOPATH_and_Modules)