package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

func main() {

	arr := [5]int{10, 20, 30, 40, 50}

	//uintptr1 := uintptr(unsafe.Pointer(&arr[0])) // 1 and 4

	//uintptr1 += 8                           //? uintptr
	//val := (*int)(unsafe.Pointer(uintptr1)) // 3 and 2
	//println(*val)
	uintptr1 := uintptr(unsafe.Pointer(&arr[0]))
	for i := 0; i < len(arr); i++ {
		val := (*int)(unsafe.Pointer(uintptr1)) // 3 and 2
		println(*val)
		uintptr1 += unsafe.Sizeof(arr[0])
	}

	str := "Hello World"
	str2 := "Hello omneNEXT"

	println("Len of str:", len(str))

	strPtr := (*[2]int)(unsafe.Pointer(&str))
	strPtr1 := (*[2]int)(unsafe.Pointer(&str2))
	strPtr[0] = strPtr1[0]
	(*strPtr)[1] = 50
	println("Len of str:", len(str))
	fmt.Println(str)
	//ptr := &arr[0]
	// fmt.Printf("0x%x\n", ptr)
	// fmt.Printf("%d\n", ptr)
	// fmt.Printf("%b\n", ptr)

	//var ptrnum int = 824633787616

	//ptrnum += 8

	// uintptr
	// pointer arithmetic --> no possible directly in Go

	// unsafe.Pointer
	// 1.A pointer value of any type can be converted to a Pointer.
	// 2.A Pointer can be converted to a pointer value of any type.
	// 3.A uintptr can be converted to a Pointer.
	// 4.A Pointer can be converted to a uintptr.

	slice1 := make([]int, 5, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = rand.IntN(50)
	}

	slicePtrs := (*[3]int)(unsafe.Pointer(&slice1))
	slicePtrs[1] = 10
	fmt.Println(slice1)

	arr1 := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	slicePtrs[0] = int(uintptr(unsafe.Pointer(&arr1[0])))
	fmt.Println(slice1)

	// var ptrarr *[3]int
	// var arrptr [3]*int
	str4 := "d2a84f4b8b650937ec8f73cd8be2c74add5a911ba64df27458ed8229da804a26"
	fmt.Println(str4[0 : len(str4)/2])
}

// void pointers

// slice:= make([]int,5,10)
// fill the slice

// 0 ptr
// 1 len --> 10
// 2 cap

// print the slice
//
