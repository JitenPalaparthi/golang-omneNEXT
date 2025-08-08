package main

import "sync"

func main() {
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		sum := GenPrimeNumsSum(100)
		println(sum)
	}(wg)
	wg.Add(1)
	sum := GenPrimeNumsSum1(wg, 10)
	println(sum)
}

func GenPrimeNumsSum(count uint) int {
	sum := 0
	if count < 3 {
		sum = 1 + 2
		return sum
	}

	for i := 3; i <= int(count); i++ {
		cp := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				cp++
			}
			if cp > 1 {
				break
			}
		}
		//println(cp)
		if cp < 1 {
			//	println(i)
			sum += i
		}
	}

	return sum
}

// This is not a good practice to use a wg as a parameter to the function or method
func GenPrimeNumsSum1(wg *sync.WaitGroup, count uint) int {
	sum := 0
	if count < 3 {
		sum = 1 + 2
		return sum
	}

	for i := 3; i <= int(count); i++ {
		cp := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				cp++
			}
			if cp > 1 {
				break
			}
		}
		//println(cp)
		if cp < 1 {
			//	println(i)
			sum += i
		}
	}
	wg.Done()
	return sum
}
