package main

const PI = 3.14

type Circle float32

func (r Circle) Area() float64 {
	return float64(PI * r * r)
}

func (r Circle) Perimeter() float64 {
	return float64(2 * PI * r)
}

func (Circle) What() string {
	return "Circle"
}
