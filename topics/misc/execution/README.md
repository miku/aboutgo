# Entry point and execution model

The entry point to a runnable Go program is a `main` function, located in a package `main`.

Before a package (main) is executed, all imported packages are initialized.
Packages are initialized in sequence. First, package level variables are
executed in declaration order. Then the `init` function runs, if defined.
Finally, main executes.

When the `main` function returns, the program exits.

## Package initialization

> Within a package, package-level variable initialization proceeds stepwise,
> with each step selecting the variable earliest in declaration order which has
> no dependencies on uninitialized variables.

More on initialization: [Package initialization](https://golang.org/ref/spec#Package_initialization)

