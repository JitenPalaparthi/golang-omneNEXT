package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type Integer = int // this is an alias

const PI float32 = 3.14

const SEC = 60

const MINUTE = 60

var Global2 int // BSS segment part of data segment

const HOUR = (SEC * MINUTE * 60) * 24 / PI

var Global1 int = 99999 // static

func main() {

	fmt.Println("Hello World", time.Now())

	var num1 int32 = 12312

	var (
		num2   int16  = 12312
		num3   uint32 = 123213321
		num4   uint64 = 123213123123123
		num5   int
		float1 float32 = 1231231.12321
		float2 float64 = 9797933.343434434
		float3 float64
		ok1    bool
		str1   string = "Hello OmneNEXT"
		str2   string
	)

	fmt.Println("Str1:", str1, "Size of str1", unsafe.Sizeof(str1))
	fmt.Println("Str2:", str2, "Size of str2", unsafe.Sizeof(str2))

	fmt.Println(num1, num2, num3, num4, num5, float1, float2, float3, ok1, str1, str2)

	var num6 = 1212423213 // int

	var (
		float4 = 123213.12312 // float64
		float5 = 54545.454    // float64
		float6 = 1.1          // float64
		age    = 40
	)

	var ok2 = true
	var str3 = "Hello World"

	fmt.Println("Value of float6:", float6, "Type of float6:", reflect.TypeOf(float6))
	fmt.Println("Value of age:", age, "Type of age:", reflect.TypeOf(age))

	fmt.Println(num6, float4, float5, ok2, str3)

	var char1 rune = 'A' // 4 bytes

	var char2 rune = 95 // 4 bytes

	var char3 int32 = 'B'

	var char4 int64 = 'A'

	var char6 rune = '好'

	var char7 uint8 = 'Z'

	var char8 uint16 = '你'

	//你  好

	var char5 = char1 + char2 + char3

	println(char1, char2, char3, char4, char5, char6, char7, char8)

	var i1 Integer = 12321312

	var byte1 byte = 200

	println(i1, byte1)

}

// runtime.main

// numbers
// uint, int, uint8,uint16,uint32,uint64,int8,int16,int32,int64, float32,float64,rune, byte

// string --> string

// bool --> bool

// any --> interface{} or any
