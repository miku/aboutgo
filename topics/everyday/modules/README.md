# Working with Go modules

> A module is a collection of related Go packages. Modules are the unit of
> source code interchange and versioning. The go command has direct support for
> working with modules, including recording and resolving dependencies on other
> modules. Modules replace the old GOPATH-based approach to specifying which
> source files are used in a given build.

If you start a new projects as of 2019, use Go modules.

> Go 1.13 includes support for Go modules. Module-aware mode is active by default whenever a go.mod file is found in, or in a parent of, the current directory.

> The quickest way to take advantage of module support is to check out your
> repository, create a go.mod file (described in the next section) there, and
> run go commands from within that file tree.

Further information on publishing packages can be found in the wiki:

* [wiki/PackagePublishing](https://github.com/golang/go/wiki/PackagePublishing)