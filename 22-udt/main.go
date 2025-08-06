package main

import "fmt"

func main() {
	greet()

	r1 := Rect{L: 43.54, B: 65.56}
	//a1, p1 := Area(r1), Area(r1)
	a1, p1 := Area(&r1), Area(&r1)

	fmt.Printf("Area of r1:%.3f Perimeter of r1: %.3f\n", a1, p1)
	fmt.Printf("Area of r1:%.3f Perimeter of r1: %.3f\n", r1.A, r1.P)

	r2 := Rect{L: 10.34, B: 12.45}
	a2, p2 := (&r2).Area(), r2.Perimeter()
	fmt.Printf("Area of r2:%.3f Perimeter of r2: %.3f\n", a2, p2)
	fmt.Printf("Area of r2:%.3f Perimeter of r2: %.3f\n", r2.A, r2.P)

}

type Rect struct {
	L, B float32
	A, P float64
}

// func Area(r Rect) float64 {
// 	r.A = float64(r.L * r.B)
// 	return r.A
// }

// func Perimeter(r Rect) float64 {
// 	r.P = float64(2 * (r.L + r.B))
// 	return r.P
// }

func Area(r *Rect) float64 {
	r.A = float64(r.L * r.B)
	return r.A
}

func Perimeter(r *Rect) float64 {
	r.P = float64(2 * (r.L + r.B))
	return r.P
}

func (r *Rect) Area() float64 {
	r.A = float64(r.L * r.B)
	return r.A
}

func (r *Rect) Perimeter() float64 {
	r.P = float64(2 * (r.L + r.B))
	return r.P
}
