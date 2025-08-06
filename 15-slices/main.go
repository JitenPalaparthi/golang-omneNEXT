package main

import "fmt"

func main() {

	slice1 := []int{10, 20, 30, 40, 50}

	slice1[0] = 100
	slice1[1] = 200
	//slice1[5] = 600

	slice1 = append(slice1, 500)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	println(len(slice))
	slice = append(slice[:4], slice[5:]...)
	fmt.Println(slice)
	println(len(slice))

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice3 := arr[3:9]

	// append slice3 to take 1 2 from arrany and prepend it
}
