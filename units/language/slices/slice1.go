package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	var s0 []string
	s1 := []string{}
	s2 := []string{"Hello", "World"}
	s3 := make([]string, 0)
	s4 := make([]string, 0, 0)
	s5 := make([]string, 0, 10)
	s6 := make([]string, 10, 10)

	w := tabwriter.NewWriter(os.Stdout, 1, 4, 2, ' ', tabwriter.Debug)

	fmt.Fprintf(w, "Value\tNil?\tType\tLength\tCapacity\n")
	for _, s := range [][]string{s0, s1, s2, s3, s4, s5, s6} {
		fmt.Fprintf(w, "%v\t%v\t%T\t%d\t%d\n", s, s == nil, s, len(s), cap(s))
	}

	w.Flush()
}
