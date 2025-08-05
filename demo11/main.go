package main

import "fmt"

func main() {

	var arr1 [5]int // zero value

	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{10, 20, 30, 40, 50} // compiler

	arr4 := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} // compiler

	arr5 := [...]string{"Hello", "World"}
	fmt.Println(arr5)
	fmt.Printf("arr1:%T arr2: %T arr3:%T arr4:%T", arr1, arr2, arr3, arr4)
	s1 := SumOf(arr1)
	s2 := SumOf(arr2)
	s3 := SumOf(arr3)
	println(s1, s2, s3)
	//s4 := SumOf(arr4)

	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	arr1[3] = 400
	arr1[4] = 500

	fmt.Println(arr1)

	var arr6 [5]any

	arr6[0] = arr1
	arr6[1] = true
	arr6[2] = "hello WOrld"
	arr6[3] = arr5
	arr6[4] = SumOf
	executeArrAny(arr6)

	arr2d := [2][2]int{{1, 2}, {3, 4}}
	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	for _, v := range arr2d {
		for _, v1 := range v {
			print(v1, " ")
		}
		println()
	}
	println()

	for _, arr1 := range arr3d {
		for _, arr2 := range arr1 {
			for _, v := range arr2 {
				print(v, " ")
			}
			println()
		}
	}

	arr7 := [5]int{10, 13, 14, 15, 16}

	var arr8 [3]int = ([3]int)(arr7)

}

func executeArrAny(any1 [5]any) {
	var arr [5]int
	for _, v := range any1 {
		switch vt := v.(type) {
		case [5]int:
			arr = vt
			fmt.Println("[5]int type array")
			for _, v1 := range vt {
				print(v1, " ")
			}
			println()
		case bool:
			fmt.Println("Bool type", vt)

		case string:
			fmt.Println("String type:", vt)

		case [2]string:
			fmt.Println("[2]string type")
			for _, v1 := range vt {
				print(v1, " ")
			}
			println()
		case func([5]int) int:
			fmt.Println("func([5]int) int")
			s := vt(arr)
			fmt.Println("Sum Of", s)
		default:
			println("unknown type")
		}

	}
}

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
