# Why Go?

## Projects describing motivations

> What happened to Python Sops? We rewrote Sops in Go to solve a number of
> deployment issues, but the Python branch still exists under python-sops. We
> will keep maintaining it for a while, and you can still pip install sops, but
> we strongly recommend you use the Go version instead.

-- [https://github.com/mozilla/sops](https://github.com/mozilla/sops)


## Performance

Go is garbage-collected, yet fast - or fast enough for many use cases.

Anecdotal evidence: A simple, unoptimized reservoir sampling tool in Go is about
2-3 times slower than a classic C tool.

-- [miku/rsampling](https://github.com/miku/rsampling)

## Not a research language

> In short, development at Google is big, can be slow, and is often clumsy. But
> it is effective.

> The goals of the Go project were to eliminate the slowness and clumsiness of
software development at Google, and thereby to make the process more productive
and scalable. The language was designed by and for people who write—and read and
debug and maintain—large software systems.

> Go's purpose is therefore not to do research into programming language design;
> it is to improve the working environment for its designers and their
> coworkers. Go is more about software engineering than programming language
> research. Or to rephrase, it is about language design in the service of
> software engineering.

-- [https://talks.golang.org/2012/splash.article](https://talks.golang.org/2012/splash.article)

