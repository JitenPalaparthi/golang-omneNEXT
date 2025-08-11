package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	//var wg *sync.WaitGroup = &sync.WaitGroup{}
	ch := make(chan int)
	result := make(chan int, 20)
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
	}(5)
	wg.Wait()
	wg1 := new(sync.WaitGroup)
	wg1.Add(1)
	close(result)
	go Receiver(result, wg1)
	wg1.Wait()

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

func Receiver(result chan int, wg *sync.WaitGroup) {
	for r := range result {
		println(r)
	}
	wg.Done()
}
