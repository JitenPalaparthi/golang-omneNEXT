package main

import "fmt"

func main() {

	var num1 uint8 = 100

	var num2 uint16 = uint16(num1)

	var num3 uint32 = 12312

	var num4 int = -5454

	var num5 uint64 = 123343

	var num6 int64 = 34234

	var sum int64 = int64(num1) + int64(num2) + int64(num3) + int64(num4) + int64(num5) + num6
	println("Sum:", sum)

	// BigInt

	var i1 uint32 = 21321323
	var i2 uint8 = uint8(i1) // 10100010101010110 01101011
	var i3 uint16 = uint16(i1)
	println(i2, i3) // 01101011

	fmt.Printf("%b", i1)

	// shorthand declare

	float1 := 13231.123

	var i4 int32 = int32(float1)

	ok1 := true

	str1 := "Hello World"

	println(float1, ok1, str1, i4)
	//var i6 uint8 = 123
	var any1 any = 123 // Value ? Type of Any1

	var i5 uint8 = uint8(any1.(int))

	println(i5)

	// var float2 float32 = any1.(float32)

	// println(float2)
	any1 = float32(12312.123)
	f1, ok := any1.(float32)

	if !ok {
		println("cannot be asserted to float32")
	} else {
		fmt.Println(f1)
	}

	any1 = 0

	i10, ok := any1.(int)

	if !ok {
		println("cannot be asserted")
	} else {
		println(i10)
	}

	a, b, c := 10, 20, 30
	// t := a
	// a = b
	// b = t

	a, b, c = b, c, a
	println(a, b, c)

	add1, sub1, mul1, div1, mod1 := calc1(10, 20)
	println(add1, sub1, mul1, div1, mod1)

	add2, sub2, _, _, _ := calc1(10, 20)
	println(add2, sub2)

	_, sub3, _, div3, _ := calc1(10, 20)
	println(sub3, div3)

}

func calc1(a, b int) (int, int, int, int, int) {
	return a + b, a - b, a * b, a / b, a % b
}

func calc2(a int, b int) (s int, su int, m int, d int) {
	s, su, m, d = a+b, a-b, a*b, a/b
	return s, su, m, d
}
