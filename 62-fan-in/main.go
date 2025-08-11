package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	gen1 := Generator("Gen-1", 10)
	gen2 := Generator("Gen-2", 20)
	gen3 := Generator("Gen-3", 30)
	gen4 := Generator("Gen-4", 5)
	gen5 := Generator("Gen-5", 15)
	gen6 := GenerateByTime("Gen-6", time.Second*5, time.Millisecond*100)
	wg.Add(1)
	go Receceiver(wg, gen1, gen2, gen3, gen4, gen5, gen6)

}

func Generator(name string, r uint) chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i < int(r); i++ {
			ch <- fmt.Sprint("Generator ", name, "-->", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
	return ch
}

func GenerateByTime(name string, duration time.Duration, delay time.Duration) chan string {
	ch := make(chan string)
	//aft := time.After(duration)
	aft := After(duration)
	go func() {
		i := 1
		done := false
		for {
			go func() {
				<-aft
				done = true
			}()
			if done {
				close(ch)
				runtime.Goexit()
			}
			ch <- fmt.Sprint("Generator ", name, "-->", i)
			time.Sleep(delay)
			i++
		}
	}()
	return ch
}

func Receceiver(wg *sync.WaitGroup, chs ...chan string) {
	for _, ch := range chs {
		wg.Add(1)
		go func(ch chan string) {
			for v := range ch {
				println(v)
			}
			wg.Done()
		}(ch)
	}
	wg.Done()
}

func After(duration time.Duration) chan struct{} {
	aft := make(chan struct{})
	go func() {
		time.Sleep(duration)
		aft <- struct{}{}
		close(aft)
	}()
	return aft
}
