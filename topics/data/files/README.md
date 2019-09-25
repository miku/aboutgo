# Working with files

## Opening files

```go
f, err := os.Open("README")
if err != nil {
    log.Fatal(err)
}
defer f.Close()
```

Files implement a couple of interfaces, e.g. `io.Reader`, `io.Writer` and a few more.

## Buffered IO

The [bufio](https://golang.org/pkg/bufio/) package offers a
[bufio.Reader](https://golang.org/pkg/bufio/#Reader), which is more efficient
and versatile.

If has great support for working with delimited data (e.g. lines), via:

* [ReadBytes](https://golang.org/pkg/bufio/#Reader.ReadBytes)
* [ReadString](https://golang.org/pkg/bufio/#Reader.ReadString)

## Temporary files and utilities

Support for temporary files and utilities can be found in the
[ioutil](https://golang.org/pkg/io/ioutil/) package.

A practical function is:
[ioutil.ReadAll](https://golang.org/pkg/io/ioutil/#ReadAll) which allows to
drain data `io.Reader` into a byte slice. However, sometimes a streaming approach can be more efficient.

## Stream processing

The io interfaces allow build pipelines.

```go
// Imagine you get a urgent request to analyze some image data. It's compressed.
// You need to find the distribution of the "red" values in an image and create a
// report in form of a pretty table.
//
// OUTPUT:
//
//     $ cat gopherbw.png.gz | go run main.go | sort -nr | head -10
//      8543296|0
//      6353501|65535
//         1346|5140
//          881|21588
//          789|5397
//          751|14135
//          677|21331
//          607|34181
//          547|11822
//          506|45489
//
// Or via web:
//
//     $ curl -s https://git.io/JesAY | go run main.go | sort -nr | head -10
//
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	// Read compressed data from standard input.
	r, _ := gzip.NewReader(os.Stdin)
	// Decode image from stream.
	img, _, _ := image.Decode(r)

	// Get image dimensions.
	rectangle := img.Bounds()
	width, height := rectangle.Size().X, rectangle.Size().Y

	// Store distribution of "red" values in a map.
	rdist := make(map[uint32]int)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, _, _, _ := img.At(x, y).RGBA()
			rdist[r]++
		}
	}

	// Write tabulated values into a buffer.
	var buf bytes.Buffer
	for k, v := range rdist {
		fmt.Fprintf(&buf, "%d\t%d\n", v, k)
	}

	// Create a pretty printer, connect it with standard output.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	// Copy buffer content to the pretty printer.
	if _, err := io.Copy(w, &buf); err != nil {
		log.Fatal(err)
	}
}
```
