package main

import "math/rand/v2"

func main() {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	a, b := 1, 1
	counter := 1
	println()
	// lika a while loop
	for counter <= 10 {
		print(a, " ")
		a, b = b, a+b
		counter++
	}

	a, b = 1, 1
	counter = 1
	println()
	// lika a while loop
	for {
		print(a, " ")
		a, b = b, a+b

		if counter >= 10 {
			break
		}
		counter++
	}
	println()
	i := 1
	for ; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()
	for i, j := 1, 10; i <= j; i, j = i+1, j-1 {
		println("i", i, "j", j)
	}
	println("nested loops")
outer:
	//done := false
	for i := 1; i <= 5; i++ {
		// if done {
		// 	break
		// }
		for j := 2; j <= 6; j++ {
			if i == j {
				//done = true
				break outer
			}
			println("i", i, "j", j)
		}
	}

	println("Rand numbers")

exit:
	for {
		num := rand.IntN(1000)
		switch {
		case num%5 == 0:
			println(num, "is multiple of 5")
			break exit
		case num%2 == 0:
			println(num, "is multiple of 2")
		case num%7 == 0:
			println(num, "is multiple of 7")
		}
	}
	count := 1
loop:
	goto printit
printit:
	println(count)
	count++
	if count <= 10 {
		goto loop
	} else {
		goto exitit
	}
exitit:
	println("Done")

}
