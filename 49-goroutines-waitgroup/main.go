package main

import (
	"runtime"
	"sync"
	"time"
)

var wg *sync.WaitGroup

func init() {
	println("init--1")
	wg = new(sync.WaitGroup)
}

func init() {
	println("init--2")
}

func init() {
	println("init--3")
}

func main() {
	wg.Add(1)
	go func() {
		println("Hello OmneNEXT")
		wg.Add(1)
		go func() {
			wg.Add(1)
			go func() {
				println("Just to check as a goroutine")
				wg.Done()
			}()
			wg.Done()
		}()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		println("Hello World")
		wg.Add(1)
		go SayHi(wg)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		c := 0
		for {
			c++
			time.Sleep(time.Second * 1)
			println("----->", c)
			if c > 10 {
				println("Exiting this gorountine")
				wg.Done()
				runtime.Goexit()
				println("after Exiting this gorountine")

			}
		}

	}()

	println("End of main")
	wg.Wait()
}

func SayHi(wg *sync.WaitGroup) {
	println("Hello Everyone")
	wg.Done()
}

// 1. main is also a goroutine
// 2. no goroutine waits for other gorouitne to complet its execution, which is including main
// 3. the order of execution is not guaranteed

// Two functions//
// Generate even numbers (name,from, to ) GenEven("GenEven-1",10,20)
// Generate Odd numbers.. (name,from, to ) GenOdd("GenOdd-1",10,20)

// Call this functions 5 times GenEven and 5 Time GenOdd

// Run them as goroutes..
// Use wait group to complete its execution
// main should be gracefully exited
