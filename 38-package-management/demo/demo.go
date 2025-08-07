package demo

import (
	"fmt"
)

var Global int

var global int

func Greet() {
	//shapes.Greet()
	fmt.Println("Greetings from demo package")
}

type T1 struct {
	A int
	b int
	C int
	d int
}

func (t *T1) print() {
	fmt.Println(t)
}

func (t *T1) Print() {
	fmt.Println(t)
}

func (t *T1) SetB(b int) {
	t.b = b
}

func (t *T1) SetD(d int) {
	t.d = d
}

// less likely written

type t1 struct {
	A int
	b int
	C int
	d int
}

func (t *t1) print() {
	fmt.Println(t)
}

func (t *t1) Print() {
	fmt.Println(t)
}
