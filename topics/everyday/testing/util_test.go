package testing

import "testing"

// Exercise: Write a function to test the MinInt function. Consider a table-driven test.

func TestMinInt(t *testing.T) {
	var cases = []struct {
		a, b int
		out  int
	}{
		{a: 0, b: 0, out: 0},
		{a: 1, b: 0, out: 1},
	}

	for _, c := range cases {
		if out := MinInt(c.a, c.b); out != c.out {
			t.Errorf("got %v, want %v", out, c.out)
		}
	}
}
