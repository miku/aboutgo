# Importing code

With Go modules, integrating third party code has become simpler.

* start your project
* run `go mod init`
* run `go get` to fetch required code, e.g. `go get github.com/fatih/color`
* keep `go.mod` and `go.sum` under version control

If you code has dependencies and it compiles, you can have reproducible builds.

## Updating all dependencies

Running:

```
$ go get -u -v ./...
```

will update all the dependencies of your project. If there are updates, the
`go.mod` and `go.sum` files will be updates and should be committed.

You can also just update a single dependency:

```
$ go get -u github.com/fatih/color
```

## Versions

You will see two kinds of entries in your `go.mod` files:

```
...
github.com/gorilla/websocket v1.4.0
github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d
...
```

The first uses git tags, the second one uses a timestamp-commit string (if no
tag is available).

## Compatibility

Full compatibility cannot be guaranteed by tools, but there are a few rules.

### Import compatibility rule

> If an old package and a new package have the same import path, the new package
> must be backwards compatible with the old package.

So the way to go about it is like this:

> By including the major version into the import path, it should be clearer to
> authors that `my/thing` and `my/thing/v2` are different and need to be able to
> coexist.

This simple pattern allows for things like gradual code repair, since two
versions of one package might coexist.

It is more work for library authors, as they need to keep in mind the potential
dependents of their packages.

### Minimal version selection

> Minimal version selection assumes that each module declares its own dependency
> requirements: a list of minimum versions of other modules.

> Modules are assumed to follow the import compatibility rule—packages in any
> newer version should work as well as older ones—so a dependency requirement
> gives only a minimum version, never a maximum version or a list of
> incompatible later versions.

In one sentence: *In contrast, minimal version selection prefers the minimum
allowed version, which is the exact version requested by some go.mod in the
project.*
