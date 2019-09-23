# Testing Go Code

Here, we focus on unit tests.

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
