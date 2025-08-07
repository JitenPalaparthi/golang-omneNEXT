package main

type Square float32

func (s Square) Area() float64 {
	return float64(s * s)
}

func (s Square) Perimeter() float64 {
	return float64(s * 4)
}

func (Square) What() string {
	return "Square"
}
