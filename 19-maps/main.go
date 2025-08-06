package main

import "fmt"

func main() {
	m1 := make(map[string]any)
	m1["add"] = add
	m1["sub"] = sub
	m1["mul"] = mul
	m1["div"] = div
	m1["greet"] = greet
	m1["sq"] = sq

	a, b, n := 10, 20, 5

	for k, v := range m1 {
		switch vf := v.(type) {
		case func(int, int) int:
			switch k {
			case "add":
				println("Add", vf(a, b))
			case "sub":
				println("Sub", vf(a, b))
			case "mul":
				println("Mul", vf(a, b))
			case "div":
				println("Div", vf(a, b))
			}
		case func():
			println("called greet")
			//v.(func())()
			vf()
		case func(int) int:
			println("Sq")
			println("Sq:", vf(n))
		default:
			println("not found")
		}

	}

	var m2 map[string]string = map[string]string{"name": "Jiten", "address": "Trv"}

	fmt.Println(m2)

	m3 := make(map[string]func(int, int) int)

	arr1 := [3]int{10, 20, 30}
	arr2 := [3]int{40, 50, 60}
	if arr1 == arr2 {

	}

	// slice1 := []int{10, 20, 30}
	// slice2 := []int{40, 50, 60}
	// if slice1 == slice2 {

	// }

	m3["add"] = add
	m3["sub"] = sub
	m3["mul"] = mul
	m3["div"] = div

	var fn func(int, int) int
	fn = add
	r := fn(100, 200)
	println(r)

	m4 := make(map[[3]int]string)
	m4[arr1] = "Some array with 3 elements"
	m4[arr2] = "another array"
}

func add(a, b int) int { // func(int,int)int
	return a + b
}

func sub(a, b int) int { // func(int,int)int
	return a - b
}

func mul(a, b int) int { // func(int,int)int
	return a * b
}

func div(a, b int) int { // func(int,int)int
	return a / b
}

func greet() { // func()
	println("Hello omneNEXT")
}

func sq(num int) int { // func(int)int
	return num * num
}
