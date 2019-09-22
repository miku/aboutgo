# Where does Go come from?

Go is an experiment. It tries to join various approaches to programming.

* https://golang.org/doc/faq#ancestors

> Go is mostly in the C family (basic syntax), with significant input from the
> Pascal/Modula/Oberon family (declarations, packages), plus some ideas from
> languages inspired by Tony Hoare's CSP, such as Newsqueak and Limbo
> (concurrency). However, it is a new language across the board. In every
> respect the language was designed by thinking about what programmers do and
> how to make programming, at least the kind of programming we do, more
> effective, which means more fun. 

## Google scale

Scaling programming to many machines and many developers - of all experience levels.

> A motivating factor for Golang was how forbidding and painful Google's C++
> build process was. A lot of Golang makes more sense if you look at it through
> this lens: above all else, don't be like C++. (https://news.ycombinator.com/item?id=9711701).

## Dynamic and Static

Dynamic and static languages have been seen as making opposite tradeoffs. Go
compiles fast, so it feels dynamic, but is statically typed with all the
advantages that comes with compile-time type checking.

## C and Pascal

The authors bring together language design elements from C (Thompson) and Pascal
(Griesemer) and also research on Newsqueak (Pike), a concurrent programming
language for writing application software with interactive graphical user
interfaces, developed in the early 1980s at Bell Labs).

## Concurrency model

Communicating sequential processes (1978).

> [...] the component processors must be able  to  communicate and to
> synchronize with  each other. Many methods of achieving this have been
> proposed. (https://www.cs.cmu.edu/~crary/819-f09/Hoare78.pdf)