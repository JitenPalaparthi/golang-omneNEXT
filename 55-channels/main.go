package main

import (
	"time"
)

type Signal struct{}

func main() {

	ch := make(chan int)
	//sig := make(chan Signal)

	sig := make(chan struct{})

	go sender(ch)
	go receiver(ch, sig)

	// 	<-sig
	// }

	// func sender(ch chan<- int) {
	// 	for i := 1; i <= 10; i++ {

	// 		ch <- i
	// 		println("sending -->", i)
	// 	}
	// }

	// func receiver(ch <-chan int, sig chan<- Signal) {
	// 	for i := 1; i <= 10; i++ {
	// 		time.Sleep(time.Millisecond * 200)
	// 		println(<-ch)
	// 	}

	// 	sig <- Signal{}
	// }

	<-sig
}

func sender(ch chan<- int) {
	for i := 1; i <= 10; i++ {

		ch <- i
		println("sending -->", i)
	}
}

func receiver(ch <-chan int, sig chan<- struct{}) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 200)
		println(<-ch)
	}

	sig <- struct{}{}
}
