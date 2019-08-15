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

It is possible to use struct tags for CSV as well (custom example:
[annotations](https://github.com/miku/span/blob/29d3a73f26f3ff547a04bb47cd9dd9c6cd60787f/licensing/entry.go#L89-L131),
[decoder](https://github.com/miku/span/blob/29d3a73f26f3ff547a04bb47cd9dd9c6cd60787f/encoding/tsv/decoder.go#L53-L86)).

A generic packages for CSV:

* [jszwec/csvutil](https://github.com/jszwec/csvutil)
* [gocarina/gocsv](https://github.com/gocarina/gocsv)
