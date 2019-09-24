# The go get command

The `go` command can fetch Go projects from the web.

> Get downloads the packages named by the import paths, along with their
> dependencies. It then installs the named packages, like 'go install'.

The semantic of `go get` depends on where the command is run. If it's run outside of a project with a go module, it behaves as describe in [Legacy GOPATH go get](https://golang.org/cmd/go/#hdr-Legacy_GOPATH_go_get).


You can host code anywhere with a [few
steps](https://jve.linuxwall.info/blog/index.php?post/2015/08/26/Hosting_Go_code_on_Github_with_custom_import_path),
but for the major hosting sites, [support is
built-in](https://github.com/golang/tools/blob/c006dc79eb54fe1d2cfcfbe2cf24783e111c368a/go/vcs/vcs.go#L649-L691).

> When choosing a domain, keep in mind that it will be the name of your package
> for the foreseeable future, so choose a name that you’ll still like tomorrow.
> (https://sagikazarmark.hu/blog/vanity-import-paths-in-go/)