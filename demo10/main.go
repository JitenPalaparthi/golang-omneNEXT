package main

func main() {

	str1 := "Hello World"
	str2 := "Hello 你好"

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), " ")
	}
	println()

	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), " ")
	}
	println()
	//str3 := ""
	for i, v := range str2 {
		println("index:", i, "-->", string(v), " ")
		//str3 := str3 + string(v)
	}
}

// for range loop

// Take a lower character string
// "hello world"
// concert it to uppercase and store it in another string
