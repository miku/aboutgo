package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b := a[2:6]
	c := a[4:]
	d := a[:]
	e := c[1:]

	expressions := []string{
		"a := []int{0, 1, 2, 3, 4, 5, 6, 7}",
		"b := a[2:6]",
		"c := a[4:]",
		"d := a[:]",
		"e := c[1:]",
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', tabwriter.Debug)

	fmt.Fprintf(w, "Expression\tValue\tNil?\tType\tLength\tCapacity\n")
	for i, s := range [][]int{a, b, c, d, e} {
		fmt.Fprintf(w, "%s\t%v\t%v\t%T\t%d\t%d\n", expressions[i], s, s == nil, s, len(s), cap(s))
	}

	w.Flush()

}
