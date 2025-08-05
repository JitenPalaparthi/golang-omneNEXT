package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []int

	if slice1 == nil {
		println("slice1 is a nil slice")
		slice1 = make([]int, 5, 5)
	}

	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	for i := range slice1 {
		slice1[i] = rand.IntN(100)
	}
	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	slice1 = append(slice1, 1111)
	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	slice1 = append(slice1, 2222)
	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	slice1 = append(slice1, 3333, 4444, 5555)
	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)

	slice1 = append(slice1, 6666)
	fmt.Printf("slice1:%v\nlen: %d\ncap: %d\nptr:%p\naddress:%p\n", slice1, len(slice1), cap(slice1), &slice1[0], &slice1)
	println("Sum:", SumOfV())
	println("Sum:", SumOfV(10, 20))
	println("Sum:", SumOfV(12, 32, 34, 34, 4, 57, 57, 8, 7))
	println("Sum:", SumOfV(slice1...))
	//fmt.Println(true, 0, "helo", 12.213, "world", 1, 2, 3, false, true, false)

	arr := [3]int{10, 20, 30}
	fmt.Println(SumOf(arr[:]))
	fmt.Println(SumOf(slice1))
}

// variadic params..
// variadia param can only be used in function or method
// variadic param must be the last param
func SumOfV(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}
	return sum
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
