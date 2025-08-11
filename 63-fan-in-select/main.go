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
	gen2 := GenerateByTime("Gen-2", time.Second*5, time.Millisecond*100)

	aft := After(time.Second * 5)
	wg.Add(1)
	go RececeiverBySelect(wg, gen1, gen2, aft)
	gen3 := Generator("Gen-3", 10)
	gen4 := GenerateByTime("Gen-4", time.Second*5, time.Millisecond*100)
	wg.Add(1)
	go RececeiverBySelectClose(wg, gen3, gen4)

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

func RececeiverBySelect(wg *sync.WaitGroup, ch1, ch2 chan string, chDuration chan struct{}) {
out:
	for {
		select {
		case v, ok := <-ch1:
			if ok {
				println(v, "-->", ok)
			}
		case v, ok := <-ch2:
			if ok {
				println(v, "-->", ok)
			}
		case <-chDuration:
			println("Time elapses.... ")
			break out
			//default:
			//print("---- hit ----")
		}
	}

	wg.Done()
}

func RececeiverBySelectClose(wg *sync.WaitGroup, ch1, ch2 chan string) {
	done1, done2 := false, false
	for {
		if done1 && done2 {
			break
		}
		select {
		case v, ok := <-ch1:
			if ok {
				println(v, "-->", ok)
			} else {
				done1 = true
			}
		case v, ok := <-ch2:
			if ok {
				println(v, "-->", ok)
			} else {
				done2 = true
			}
		}

	}

	wg.Done()
}
