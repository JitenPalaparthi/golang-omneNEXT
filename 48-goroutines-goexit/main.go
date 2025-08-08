package main

import (
	"runtime"
	"time"
)

func main() {
	go func() {
		println("Hello OmneNEXT")
		go func() {
			go println("Just to check as a goroutine")
		}()
	}()
	go func() {
		println("Hello World")
		go SayHi()
	}()

	go func() {
		c := 0
		for {
			c++
			time.Sleep(time.Second * 1)
			println("----->", c)
			if c > 10 {
				println("Exiting this gorountine")
				runtime.Goexit()
				println("after Exiting this gorountine")

			}
		}
	}()

	println("End of main")
	time.Sleep(time.Second * 11)
}

func SayHi() {
	println("Hello Everyone")
}

// 1. main is also a goroutine
// 2. no goroutine waits for other gorouitne to complet its execution, which is including main
// 3. the order of execution is not guaranteed
