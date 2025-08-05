package main

import (
	"math/rand/v2"
	"strconv"
)

func main() {

	age := uint8(18)

	if age >= 18 {
		println("eligible for vote")
	} else {
		println("not eligible for vote")
	}

	char := 'm'

	age = 21

	if age >= 21 && (char == 'm' || char == 0x4d) {
		println("He eligible for marriage")
	} else if age >= 18 && (char == 'f' || char == 70) {
		println("she is eligible for marraige")
	} else {
		println("not eligible")
	}

	str := "123123"

	if num, err := strconv.Atoi(str); err != nil {
		println(err.Error())
	} else {
		println("The number is ", num)
	}

	if num := rand.IntN(10000); num%2 == 0 {
		println(num, " is an even number")
	} else {
		println(num, " is a odd number")
	}

	//println(num)

}
