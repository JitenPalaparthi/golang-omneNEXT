package main

import "fmt"

func main() {

	//	slice1 := []int{1, 2, 3, 4, 5} //ptr: len:5 cap:5
	slice1 := make([]int, 5, 5)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 1, 2, 3, 4, 5
	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)

	// slice2 := slice1 // The headers are copied
	// fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)

	// slice2[0] = 9999

	// fmt.Println(slice1)

	// slice2 = append(slice2, 20)
	// slice2[1] = 8888
	// fmt.Println(slice1)
	// fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	// fmt.Printf("slice2:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice2, len(slice2), cap(slice2), &slice2[0], &slice2)
	// sq(slice1)
	// fmt.Println(slice1)
	AddSq(slice1, 6, 7, 8, 9, 10)
	fmt.Println(slice1)
	slice1 = AddSqR(slice1, 6, 7, 8, 9, 10)
	fmt.Println(slice1)
}

func sq(slice []int) {
	for i, v := range slice {
		slice[i] = v * v
	}
}

func AddSq(slice []int, nums ...int) {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
}

func AddSqR(slice []int, nums ...int) []int {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}
