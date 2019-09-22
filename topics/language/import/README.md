# Import statements

> An import declaration states that the source file containing the declaration
> depends on functionality of the imported package and enables access to
> exported identifiers of that package. The import names an identifier
> (PackageName) to be used for access and an ImportPath that specifies the
> package to be imported.

Example `Sin` function from `math` package in the standard library.

```
Import declaration          Local name of Sin

import   "math"             math.Sin
import m "math"             m.Sin
import . "math"             Sin
import _ "math"             -
```

## Factored import

Multiple import statements can be grouped.

```go
import "io"
import "bufio"
```

The above would be equivalent to:

```go
import (
    "io"
    "bufio"
)
```

## Unused import

A program with an import of a package that is not used will not compile.

Example error: [play.golang.org/p/mFROnlITE0X](https://play.golang.org/p/mFROnlITE0X)

```go
package main

import (
	"fmt"
)

func main() {}
```

> ./prog.go:4:2: imported and not used: "fmt"
