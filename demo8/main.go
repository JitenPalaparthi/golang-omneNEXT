package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	if s1, err := add(10, 20); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s1)
	}

	if s2, err := add(uint8(100), uint8(200)); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s2)
	}

	s3, err := add(76.46, 43.54)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s3)
	}

	s4, err := add(float32(76.23), float32(12.34))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		//fmt.Println(math.Round(s4))
		fmt.Printf("%.3f\n", s4)
	}

	if s5, err := add(10, 10.45); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s5)
	}

	if s6, err := add(true, false); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s6)
	}

	// using Generics

	r1 := addG(10, 120)
	fmt.Println(r1)
	r2 := addG(10.43, 120.34)
	fmt.Println(r2)

	r3 := addG(123, 123.32)
	fmt.Println(r3)

	// r4 := addG(true, false)
	// println(r4)
}

// %s %v %T %b %x %d %f for type reflect.TypeOf()

func IsNumber(n any) bool {
	// type switch
	switch n.(type) {
	case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	}
	return false
}

func add1(a int, b, c int) int {
	return a + b + c
}

func add(a any, b any) (float64, error) {
	//sum := float64(0)
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, errors.New("a and b are different types")
	}
	// sure that a abd a are same types
	if !IsNumber(a) {
		return 0, errors.New("input argument is not a number type")
	}
	switch a.(type) {
	case int:
		return float64(a.(int) + b.(int)), nil
	case uint:
		return float64(a.(uint) + b.(uint)), nil
	case uint8:
		return float64(a.(uint8) + b.(uint8)), nil
	case uint16:
		return float64(a.(uint16) + b.(uint16)), nil
	case uint32:
		return float64(a.(uint32) + b.(uint32)), nil
	case uint64:
		return float64(a.(uint64) + b.(uint64)), nil
	case int8:
		return float64(a.(int8) + b.(int8)), nil
	case int16:
		return float64(a.(int16) + b.(int16)), nil
	case int32:
		return float64(a.(int32) + b.(int32)), nil
	case int64:
		return float64(a.(int64) + b.(int64)), nil
	case float32:
		return float64(a.(float32) + b.(float32)), nil
	case float64:
		return a.(float64) + b.(float64), nil
	}

	return 0, nil
}

func addG[T uint | int | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64](a T, b T) T {
	return a + b
}

func addGS[T AddType](a T, b T) T {
	return a + b
}

type AddType interface {
	uint | int | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64
}
