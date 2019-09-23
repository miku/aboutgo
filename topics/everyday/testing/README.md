# Testing Go Code

Here, we focus on unit tests. Go provides a basic testing framework out of the box. It lacks features, for example assertions - which is partly explained [in the FAQ](https://golang.org/doc/faq#testing_framework):

> Go's standard testing package makes it easy to write unit tests, but it lacks
> features provided in other language's testing frameworks such as assertion
> functions.

## File naming convention

If you have a file `readerutil.go`, put tests into `readerutil_test.go`. The `go
test` command will find these files automatically.

> 'Go test' recompiles each package along with any files with names matching the
file pattern "*_test.go".

Test files can contain:

* Test functions
* Benchmark functions
* Example functions

## Function names

> A test function is one named `TestXxx` (where `Xxx` does not start with a lower
case letter) and should have the signature,

    func TestXxx(t *testing.T) { ... }

> A benchmark function is one named `BenchmarkXxx` and should have the signature,

    func BenchmarkXxx(b *testing.B) { ... }

The [`testing.T`](https://golang.org/pkg/testing/#T) and
[`testing.B`](https://golang.org/pkg/testing/#B) types allow to signal failure
and logging.

## Table driven tests

Go has a preference for table driven tests, utilizing anonymous structs.

* [wiki/TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)

```go
var flagtests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
	{"%+a", "[%+a]"},
	{"%#a", "[%#a]"},
	{"% a", "[% a]"},
	{"%0a", "[%0a]"},
	{"%1.2a", "[%1.2a]"},
	{"%-1.2a", "[%-1.2a]"},
	{"%+1.2a", "[%+1.2a]"},
	{"%-+1.2a", "[%+-1.2a]"},
	{"%-+1.2abc", "[%+-1.2a]bc"},
	{"%-1.2abc", "[%-1.2a]bc"},
}
func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := Sprintf(tt.in, &flagprinter)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
```

