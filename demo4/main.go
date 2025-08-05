package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	char1 := 'A'

	println(char1, string(char1))

	num1 := 22000
	println(string(num1))

	str1 := strconv.Itoa(num1)

	//str1 := fmt.Sprint(num1)

	fmt.Println(str1, "type of str1:", reflect.TypeOf(str1))

	str2 := "123a213"

	i1, err := strconv.Atoi(str2)
	if err != nil {
		println(err.Error())
	} else {
		println(i1)
	}

	i2 := 312312
	ok1 := true
	str3 := "Hello World"

	str4 := fmt.Sprint(i2, " ", ok1, " ", str3)

	//println(str4)
	fmt.Println(str4, "type of str4:", reflect.TypeOf(str4))
}

// strconv.ParseBool
// strconv.ParseFloat
// strconv.ParseComplex

// BigInt
// Complex
