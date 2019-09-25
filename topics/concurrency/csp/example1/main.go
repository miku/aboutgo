package main

import (
	"fmt"
	"runtime"
)

func main() {

	// NumCPU returns the number of logical
	// CPUs usable by the current process.
	fmt.Println(runtime.NumCPU())

	// go func() {}()
	fmt.Println(runtime.NumGoroutine())
	// time.Sleep(1 * time.Second)

}
