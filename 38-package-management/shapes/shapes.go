package shapes

import "fmt"

func PrintShape(shaper Shaper) {
	fmt.Printf("Area of %s: %0.2f  \n", shaper.What(), shaper.Area())
	fmt.Printf("Perimeter of %s: %0.2f \n", shaper.What(), shaper.Perimeter())
}

func Greet() {
	greet()
}

type Shaper interface {
	// Area() float64
	// Perimeter() float64
	IArea
	IPerimeter
	IWhat
}

type IArea interface {
	Area() float64
}

type IPerimeter interface {
	Perimeter() float64
}

type IWhat interface {
	What() string
}
