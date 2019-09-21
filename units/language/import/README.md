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
```
