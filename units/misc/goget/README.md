# The go get command

The `go` command can fetch Go projects from the web.

> Get downloads the packages named by the import paths, along with their
> dependencies. It then installs the named packages, like 'go install'.

The semantic of `go get` depends on where the command is run. If it's run outside of a project with a go module, it behaves as describe in [Legacy GOPATH go get](https://golang.org/cmd/go/#hdr-Legacy_GOPATH_go_get).


