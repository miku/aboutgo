# CSV

CSV support is part of the standard library.

> Package csv reads and writes comma-separated values (CSV) files. There are many kinds of CSV files; this package supports the format described in RFC 4180. 

```go
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
```