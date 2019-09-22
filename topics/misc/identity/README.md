# Truthiness and identity

## Truth

There is a type `bool` with two predeclared identiers: `true` and `false`. The
logical operators are `!`, `||` and `&&`.

## Type Identity

> A defined type is always different from any other type. Otherwise, two types
> are identical if their underlying type literals are structurally equivalent;
> that is, they have the same literal structure and corresponding components
> have identical types.

> https://golang.org/ref/spec#Type_identity

## Struct Identity

Can structs be compared? Yes, up to a certain point.

```go
package main

import (
	"fmt"
)

type Thing struct {
	Value int
}

type Complex struct {
	S string
	M map[string]string
}

func main() {
	a, b := Thing{1}, Thing{1}
	fmt.Println(a == b)   // true
	fmt.Println(&a == &b) // false
	fmt.Println(&a == &a) // true

	a, b = Thing{1}, Thing{2}
	fmt.Println(a == b) // false

	c, d := Complex{}, Complex{}
	fmt.Println(c == d) // invalid operation: c == d (struct containing map[string]string cannot be compared)
}
```

* [https://play.golang.org/p/bX9rzmF5Tdz](https://play.golang.org/p/bX9rzmF5Tdz)
* [How to compare struct, slice, map are equal?](https://stackoverflow.com/questions/24534072/how-to-compare-struct-slice-map-are-equal)

There are at least two more sophisticated ways to compare two values:

* [reflect.DeepEqual](https://golang.org/pkg/reflect/#DeepEqual) from the
  standard library

> DeepEqual reports whether x and y are “deeply equal,” defined as follows. Two
> values of identical type are deeply equal if one of the following cases
> applies. Values of distinct types are never deeply equal.

> Array values are deeply equal when their corresponding elements are deeply
> equal.

> Struct values are deeply equal if their corresponding fields, both exported
> and unexported, are deeply equal.

> Func values are deeply equal if both are nil; otherwise they are not deeply
> equal.

> Interface values are deeply equal if they hold deeply equal concrete values.

> Map values are deeply equal when all of the following are true: they are both
> nil or both non-nil, they have the same length, and either they are the same
> map object or their corresponding keys (matched using Go equality) map to
> deeply equal values.

> Pointer values are deeply equal if they are equal using Go's == operator or if
> they point to deeply equal values.

> Slice values are deeply equal when all of the following are true: they are
> both nil or both non-nil, they have the same length, and either they point to
> the same initial entry of the same underlying array (that is, &x[0] == &y[0])
> or their corresponding elements (up to length) are deeply equal. Note that a
> non-nil empty slice and a nil slice (for example, []byte{} and []byte(nil))
> are not deeply equal.

> Other values - numbers, bools, strings, and channels - are deeply equal if
> they are equal using Go's == operator.

* [go-cmp](https://github.com/google/go-cmp) from Google (package intended to be
  a more powerful and safer alternative to reflect.DeepEqual for comparing
  whether two values are semantically equal.)

