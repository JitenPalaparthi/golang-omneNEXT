package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
)

var (
	LineCount      int
	WordCount      int
	VovelCount     int
	ConsonentCount int
	mu             *sync.Mutex     = new(sync.Mutex)
	wg             *sync.WaitGroup = new(sync.WaitGroup)
	mapWords       map[string]int  = make(map[string]int)
	chBytes        chan []byte     = make(chan []byte)
	chLine1        chan string     = make(chan string, 10)
	chLine2        chan string     = make(chan string, 10)
	chWords        chan []string   = make(chan []string, 10)
	chWords2       chan []string   = make(chan []string, 10)

// sig            chan struct{}   = make(chan struct{})
)

func main() {

	wg.Add(1)
	go func() {
		bytes, err := os.ReadFile("sports.txt")
		if err != nil {
			println(err.Error())
			return
		}
		chBytes <- bytes
		close(chBytes)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(bytes.NewReader(<-chBytes))
		for scanner.Scan() {
			mu.Lock()
			LineCount++
			mu.Unlock()
			line := scanner.Text()
			chLine2 <- line
			chLine1 <- line
			//chWords2 <- strings.Split(line, " ")
		}
		close(chLine1)
		close(chLine2)
		//close(chWords2)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for line := range chLine1 {
			wg.Add(1)
			go func(line string) {
				if line != "\n" {
					words := strings.Split(line, " ")
					chWords <- words
					// for _, word := range words {
					// 	mu.Lock()
					// 	mapWords[word] += 1
					// 	// _, ok := mapWords[word]
					// 	// if ok {
					// 	// 	mapWords[word] += 1
					// 	// } else {
					// 	// 	mapWords[word] = 1
					// 	// }
					// 	mu.Unlock()
					// }
					mu.Lock()
					WordCount += len(words)
					mu.Unlock()
				}
				wg.Done()
			}(line)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for line := range chLine2 {
			for _, v := range line {
				switch v {
				case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
					mu.Lock()
					VovelCount++
					mu.Unlock()
				default:
					mu.Lock()
					ConsonentCount++
					mu.Unlock()
				}
			}
		}
		wg.Done()
	}()
	wg1 := new(sync.WaitGroup)
	wg1.Add(1)
	go func() {
		for words := range chWords {
			for _, word := range words {
				mu.Lock()
				mapWords[word] += 1
				// _, ok := mapWords[word]
				// if ok {
				// 	mapWords[word] += 1
				// } else {
				// 	mapWords[word] = 1
				// }
				mu.Unlock()
			}
		}
		wg1.Done()
	}()

	go func() {
		wg.Wait()
		close(chWords)
	}()
	wg1.Wait()
	println("Line Count:", LineCount)
	println("Word Count:", WordCount)
	println("Vovel Count:", VovelCount)
	println("Consonent Count:", ConsonentCount)
	for k, v := range mapWords {
		fmt.Println("Key-->", k, "Value-->", v)
	}

}
