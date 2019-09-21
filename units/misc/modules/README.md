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

