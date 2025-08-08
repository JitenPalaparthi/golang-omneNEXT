package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	ch := make(chan int)
	defer wg.Wait()

	wg.Add(2)
	go sender(wg, ch)
	go receiver(wg, ch)

}

func sender(wg *sync.WaitGroup, ch chan int) {
	for i := 1; i <= 10; i++ {

		ch <- i
		println("sending -->", i)
	}
	wg.Done()
}

func receiver(wg *sync.WaitGroup, ch chan int) {
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 200)
		println(<-ch)
	}
	wg.Done()
}
