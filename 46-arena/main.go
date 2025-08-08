package main

import (
	"arena"
	"fmt"
)

func main() {
	// Create a new arena
	a := arena.NewArena()
	defer a.Free() // Free all memory in the arena when done

	// Allocate objects in the arena
	nums := arena.MakeSlice[int](a, 0, 10000000000)
	nums = append(nums, 1, 2, 3)

	fmt.Println(nums) // Output: [1 2 3]

}

// GOEXPERIMENT=arenas go run main.go
