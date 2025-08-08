package main

import "sync"

func main() {
	wg := new(sync.WaitGroup)
	counter := New(0)
	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go counter.Increment(wg)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			go counter.Decrement(wg)
		}
		wg.Done()
	}()
	wg.Wait()
	println(counter.GetCount())

}

// Do not communicate by sharing memory; instead, share memory by communicating.

type Counter struct {
	count int
	mu    *sync.Mutex
}

func New(c int) *Counter {
	return &Counter{c, &sync.Mutex{}}
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	wg.Done()
}

func (c *Counter) Decrement(wg *sync.WaitGroup) {
	c.mu.Lock()
	c.count--
	c.mu.Unlock()
	wg.Done()
}

func (c *Counter) GetCount() int {
	return c.count
}
