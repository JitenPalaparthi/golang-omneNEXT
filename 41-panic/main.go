package main

import (
	"math/rand/v2"
	"os"
)

func main() {
	func() { // func1
		func() { // func2
			func() { //func3
				// num := 0
				// println(100 / num)
				num := rand.IntN(100)
				if num%2 == 0 {
					//panic("panic bcz div by 5") // user defined panic
					Fatalthis("something wernt wrong.. So fatal here")
				}
			}()
			println("Done-1")
		}()
		println("Done-2")
	}()

	println("Done-3")

	// var ptr *int
	// fmt.Println(*ptr)

	arr := [2]int{10, 11}
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}

func Fatalthis(msg string) {
	println(msg)
	os.Exit(1)
}
