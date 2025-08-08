package main

import "time"

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

	println("End of main")
	time.Sleep(time.Millisecond * 10)
}

func SayHi() {
	println("Hello Everyone")
}

// 1. main is also a goroutine
// 2. no goroutine waits for other gorouitne to complet its execution, which is including main
// 3. the order of execution is not guaranteed
