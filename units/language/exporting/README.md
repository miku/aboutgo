# Exporting

In Go, a name is exported if it begins with a capital letter.


```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.pi)
}
```

The above program will not compile.

```
./prog.go:9:14: cannot refer to unexported name math.pi
./prog.go:9:14: undefined: math.pi
```

Any uppercase variable name, type and function will be public, anything else
will be visible only within the package (which might contain multiple files).
