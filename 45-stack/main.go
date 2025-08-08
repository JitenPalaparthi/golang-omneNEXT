package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Allocate a buffer to hold the stack trace
	buf := make([]byte, 1024)

	// Capture the stack trace
	n := runtime.Stack(buf, true)

	// Print the stack trace
	fmt.Printf("Stack trace:\n%s", buf[:n])

}
