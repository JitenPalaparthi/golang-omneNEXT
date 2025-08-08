package main

import "sync"

var counter int = 0

func main() {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Increment(wg, mu)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go Decrement(wg, mu)
		}
		wg.Done()
	}()
	wg.Wait()
	println(counter)

}

func Increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func Decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter--
	mu.Unlock()
	wg.Done()
}

// Do not communicate by sharing memory; instead, share memory by communicating.
