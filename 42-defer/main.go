package main

import "os"

func main() { // main SF
	Write("data.txt", []byte("Hello omneNEXT"))
	defer func() { // func1
		defer println("end of main")
		defer println("end of func1")
		num := 0
		println(100 / num) // sure panic
		println("start of func1")
	}()

	println("start of main")
}

func Write(filename string, p []byte) (n int, err error) {
	f, err := os.OpenFile(filename, os.O_RDWR, 0644)

	if err != nil {
		panic("file is not available")
		//return 0, err
	}
	defer f.Close()
	// n, err = f.Write(p)
	// return n, err

	return f.Write(p)
}
