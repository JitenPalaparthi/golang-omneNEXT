package main

import (
	"fmt"
	"reflect"
)

func main() {

	func() {
		println("Hello omneNEXT")
	}()

	r1 := func(a, b int) int {
		return a + b
	}(10, 30)
	println("result r1:", r1)

	//var fn1 func(int, int) int
	fn1 := func(a, b int) int {
		return a * b
	}
	fn1 = sub
	var fn2 func(int, int) int

	if fn2 == nil {
		println("Function is nil")
	}

	r2 := fn1(10, 4)
	println("result r2:", r2)

	//

	r3 := calc(10, 20, func(i1, i2 int) int {
		return i1 / i2
	})

	r4 := calc(30, 20, sub)
	println(r3, r4)
	println(sub)

	slice1 := make([]func(), 5)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = func() {
			println(i)
		}
	}

	for _, fn := range slice1 {
		fn()
	}

	slice2 := make([]func(int), 5)
	i := 0
loop:
	//println(i)
	slice2[i] = func(j int) {
		println(j)
	}
	i++
	if i <= 4 {
		goto loop
	}

	for i, fn := range slice2 {
		fn(i)
	}

	sq := getSq(13)

	s := sq()
	println(s)

	execute(10, 20, func(i1, i2 int) int { return i1 + i2 }, func(i1, i2 int) int { return i1 - i2 }, func(i1, i2 int) int { return i1 * i2 }, func(i1, i2 int) int { return i1 / i2 })
	fmt.Println(reflect.TypeOf(getSq))
}

func sub(a, b int) int {
	return a - b
}

// func(int,int,func(int,int)int)int
func calc(a, b int, fn func(int, int) int) int {
	r := fn(a, b)
	return r
}

// func(int) func()int
func getSq(n int) func() int {
	println("input number is ", n)
	if n == 0 {
		n = 1
	}
	return func() int {
		return n * n
	}
}

func execute(a, b int, funcs ...func(int, int) int) {
	for _, v := range funcs {
		r := v(a, b)
		println(r)
	}
}

// func execute2(a,b int, funcs ...any){
// 4 funcs of add,sub,mul,div, getSq,cals
