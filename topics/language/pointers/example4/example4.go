// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// main is the entry point for the application.
func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2) // prefer fmt.Println, see https://golang.org/ref/spec#Bootstrapping
}

// createUserV1 creates a user value and passed
// a copy back to the caller.
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

	return u
}

// createUserV2 creates a user value and shares
// the value with the caller.
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u
}

/*
// See escape analysis and inlining decisions.

$ go build -gcflags -m=2
# github.com/miku/aboutgo/topics/language/pointers/example4
./example4.go:24:6: cannot inline createUserV1: marked go:noinline
./example4.go:38:6: cannot inline createUserV2: marked go:noinline
./example4.go:14:6: cannot inline main: non-leaf function
./example4.go:30:16: createUserV1 &u does not escape
./example4.go:46:9: &u escapes to heap
./example4.go:46:9: 	from ~r0 (return) at ./example4.go:46:2
./example4.go:39:2: moved to heap: u
./example4.go:44:16: createUserV2 &u does not escape
./example4.go:18:16: main &u1 does not escape
./example4.go:18:27: main &u2 does not escape

*/
