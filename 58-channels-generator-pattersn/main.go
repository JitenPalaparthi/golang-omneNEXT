package main

import "time"

func main() {
	ch := Generate(50)
	// sig := Receiver(ch)
	// <-sig
	<-Receiver(ch)
}

func Generate(n uint) chan int {
	ch := make(chan int)

	go func() {
		for i := 1; i <= int(n); i++ {
			time.Sleep(time.Millisecond * 100)
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func Receiver(ch chan int) chan struct{} {
	sig := make(chan struct{})
	go func() {
		for v := range ch {
			println(v)
		}
		sig <- struct{}{}
		close(sig)
	}()
	return sig
}
