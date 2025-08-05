package main

import "fmt"

func main() {

	var s1 []int // nil

	s2 := make([]int, 0)

	s3 := []int{}

	if s1 == nil {
		println("s1 is nil")
	}

	if s2 == nil {
		println("s2 is nil")
	}
	if s3 == nil {
		println("s3 is nil")
	}

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice1 := slice
	slice2 := slice[:]
	slice3 := slice[:5] // 0 to 5 but not 5
	slice4 := slice[3:8]
	slice5 := slice[5:]

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	fmt.Println("slice5", slice5)
	slice4 = append(slice4, 8888)
	fmt.Println("slice1", slice1)

	slice6 := make([]int, 10)
	copy(slice6, slice1) // deep
	fmt.Println(slice6)
	slice7 := make([]int, 3)
	copy(slice7, slice1)
	fmt.Println(slice7)

	slice8 := make([]int, 20)
	copy(slice8, slice1)
	fmt.Println(slice8)
	clear(slice1) // makes the slice to zero values
	fmt.Println(slice1)
}

// print,println,len,cap, make,append, copy, clear,
// break, case, default, else, func, if , switch, range, for, contine, fallthrough, goto, package, const, var, import,return
