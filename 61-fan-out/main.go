package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	//var wg *sync.WaitGroup = &sync.WaitGroup{}
	ch := make(chan int, 10)
	//sig := make(chan struct{})
	result := make(chan int, 2)
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

	// go Receiver(result, sig)
	// <-sig

	wg1 := new(sync.WaitGroup)
	defer wg1.Wait()

	wg1.Add(1)
	go Receiver1(result, "receiver-1", wg1)
	wg1.Add(1)
	go Receiver1(result, "receiver-2", wg1)
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

func Receiver1(result chan int, rec string, wg *sync.WaitGroup) {
	for r := range result {
		println("received by", rec, "-->", r)
	}
	wg.Done()
}
