package main

func main() {
	num := 0

	defer func() { // func1
		println("in defer-1", num)
	}()

	defer func(num int) { // func2
		println("in defer-2", num)
	}(num)

	defer println()

	num += 1
	println("normal", num)

	// {
	// 	num := 100 // new variable
	// 	{
	// 		num := 200
	// 		println(num)
	// 	}
	// 	println(num)
	// }

	str := "Hello omneNEXT"

	for _, v := range str {
		defer print(string(v))
	}

}
