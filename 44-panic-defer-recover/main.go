package main

import (
	"fmt"
)

func main() {
	println("Start of main")
	defer println("End of main")

	fn := ""
	ln := "NEXT"

	name := FullName(&fn, &ln)
	println("--->", name)
	panic("Some panic")
	func() {
		a, b := 1, 1
		for i := 1; i <= 50; i++ {
			print(a, " ")
			a, b = b, a+b
		}
	}()

}

func FullName(fn, ln *string) *string {
	ret := new(string)
	defer func() {
		println("I am called how ever it is")
		if r := recover(); r != nil {
			fmt.Println(r)
			// time.Sleep(time.Second * 1)
			// *ret = "Hello WOrld"

		}
	}()

	if fn == nil || *fn == "" {
		panic("first name cannot be empty")
	}
	if ln == nil || *ln == "" {
		panic("last name cannot be empty")
	}
	// println(*fn + *ln)
	*ret = *fn + *ln
	return ret
}
