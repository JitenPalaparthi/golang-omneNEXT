package main

import (
	"fmt"
	"math"
)

type myint1 int

func (m myint1) ToString() string {
	return fmt.Sprint(m)
}

type myint2 int

func (m myint2) Sq() myint2 {
	return m * m
}

type myint3 int

func (m myint3) Cube() myint3 {
	return m * m * m
}

type myint4 myint3

func (m myint4) Sqroot() float64 {
	return math.Sqrt(float64(m))
}

func main() {

	var i1 int = 100
	var m1 myint1 = 200
	var m2 myint2 = 300
	var m3 myint3 = 400
	var m4 myint4 = 500

	s1 := myint1(i1).ToString()
	println(s1)
	sq1 := myint2(i1).Sq()
	println(sq1)
	cb1 := myint3(i1).Cube()
	println(cb1)

	sqrt1 := myint4(i1).Sqroot()
	println(sqrt1)

	s2 := m1.ToString()
	println(s2)

	sq2 := myint2(m1).Sq()
	println(sq2)

	cb2 := myint3(m1).Cube()
	println(cb2)

	s3 := myint1(m2).ToString()
	println(s3)

	sq3 := m2.Sq()
	println(sq3)

	cb3 := myint3(m2).Cube()
	println(cb3)

	s4 := myint1(m3).ToString()
	println(s4)

	sq4 := myint2(m3).Sq()
	println(sq4)

	cb4 := m3.Cube()
	println(cb4)

	var f1 float32 = 123.123

	s5 := myint1(f1).ToString()
	println(s5)

	sq5 := myint2(f1).Sq()
	println(sq5)

	cb5 := myint3(f1).Cube()
	println(cb5)

	s6 := myint1(m4).ToString()
	println(s6)

	sq6 := myint2(m4).Sq()
	println(sq6)

	cb6 := myint3(m4).Cube()
	println(cb6)

}

// User defined number type

// any number type can be casted to any other number type

// a bool can it be casted to number? No..

// func ConvertBool(b bool) int {
// 	if b {
// 		return 1
// 	}
// 	return 0
// }
