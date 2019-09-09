package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {

	w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', tabwriter.Debug)
	fmt.Fprintf(w, "Value\tNil?\tType\tLength\tCapacity\n")

	v := [6]int{0, 1, 2, 3, 4, 5}

	fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", nil, &v == nil, v, len(v), cap(v))

	u := v[:]

	fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", u, u == nil, u, len(u), cap(u))
	u = u[:4]

	fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", u, u == nil, u, len(u), cap(u))
	u = u[:0]

	fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", u, u == nil, u, len(u), cap(u))
	// Question: Before you hit run! What do the following two lines result in?

	// u = u[1:] // -> panic
	u = u[1:3]
	// fmt.Printf("%v - len: %d, cap: %d\n", w, len(w), cap(w))

	// Theoretically.

	z := make([]int, 0, 10)
	fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", z, z == nil, z, len(z), cap(z))
	zz := z[1:2] // works, but z[1:] would not, why? z[1:] would be z[1:0], since slice len is zero.
	fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", zz, zz == nil, zz, len(zz), cap(zz))

	// j := 0 // if we dynamically set the high value, we get a panic as well

	// zz = z[1:0] // ./slice2.go:40:8: invalid slice index: 1 > 0
	// fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", zz, zz == nil, zz, len(zz), cap(zz))

	w.Flush()
}
