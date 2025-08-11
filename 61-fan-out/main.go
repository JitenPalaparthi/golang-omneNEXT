package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	//var wg *sync.WaitGroup = &sync.WaitGroup{}
	ch := make(chan int, 10)
	sig := make(chan struct{})
	result := make(chan int)
	workers := uint(5)
	wg.Add(1)
	go Publish(ch, 20, wg)

	wg.Add(1)
	go func(rec uint) {
		for i := 1; i < int(rec); i++ {
			wg.Add(1)
			go func(job int) {
				for c := range ch {
					println("received from job-->", i, "Value-->", c)
					result <- c * c
				}
				wg.Done()
			}(i)
		}
		wg.Done()
	}(workers)

	go func() {
		wg.Wait()
		close(result)
	}()

	go Receiver(result, sig)
	<-sig
}

func Publish(ch chan int, num uint, wg *sync.WaitGroup) {
	if ch == nil {
		panic("nil channel ")
	}
	for i := 1; i <= int(num); i++ {
		time.Sleep(time.Millisecond * 100)
		ch <- i
	}
	close(ch)
	wg.Done()
}

func Receiver(result chan int, sig chan struct{}) {
	for r := range result {
		println(r)
	}
	sig <- struct{}{}
}
