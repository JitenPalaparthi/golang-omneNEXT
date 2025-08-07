package main

func main() {
	// var any1 any
	// var any2 interface{}
	var e1 Empty

	e1 = 100
	e1 = "Hello WOrld"
	e1 = true

	v, ok := e1.(int)

	if ok {
		println(v)
	} else {
		println("invalid type")
	}

}

type Empty interface{}
