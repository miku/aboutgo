// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

import "fmt"

// Add imports.

// Declare a type named user.

func makeAdder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func funcName(a int) (int, error) /* (pointer return arg, error return arg) */ {

	// return a + 1, nil
	return a + 1, fmt.Errorf("some err occured")

	// Create a value of type user and return the proper values.
}

func main() {

	if _, err := funcName(10); err != nil {
		fmt.Printf("failed with %v", err)
		return
	}
	fmt.Println("ok")

	f := makeAdder(10)
	fmt.Println(f(100)) // 110

	// Use the function to create a value of type user. Check
	// the error being returned.

	// Display the value that the pointer points to.

	// Call the function again and just check the error.
}
