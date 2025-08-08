package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	var ch chan int        // only a declaration
	ch = make(chan int, 2) // unbuffered channel
	wg.Add(1)
	go func() {
		ch <- 100
		ch <- 200
		ch <- 300
		println("The value is sent")
		wg.Done()
	}()
	// blocked for 10 seconds
	time.Sleep(time.Second * 5)
	v := <-ch
	println(v)
	wg.Wait()
}

// chan is a keyword to use to create a channel
// the sender is blocked until the receiver receives the value
// the received is blocked until the sender sends the values
// To send a value to the channel ch <- 100
// <-ch to receive values from the channel
