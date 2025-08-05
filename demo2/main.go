package main

import "fmt"

func main() {

	var any1 any

	var any2 interface{} // < 1.18

	fmt.Println("Value of any1", any1)
	fmt.Println("Value of any2", any2)

	if any1 == nil {
		println("Yes any1 is nil")
	}

	any1 = 100

	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	any1 = true
	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	any1 = "Hello World"
	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	any1 = 12312.132
	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	any1 = any2
	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	var num1 = 100
	//fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	any1 = num1
	fmt.Printf("Value of any1:%v Type of any1:%T\n", any1, any1)

	if greet == nil {
		println("YEs greet is nil")
	}

	var f1 func()

	if f1 == nil {
		println("Yes f1 is nil")
	}

	f1 = greet // Contains the pointer of Greet

	f1()
}

func greet() {
	println("Hello World")
}
