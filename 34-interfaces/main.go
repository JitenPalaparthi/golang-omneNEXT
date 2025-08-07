package main

import "fmt"

func main() {

	shapesSlice := make([]Shaper, 5)
	shapesSlice[0] = NewRect(10.23, 7.89)
	shapesSlice[1] = Square(32.23)
	shapesSlice[2] = Circle(56.7)
	shapesSlice[3] = NewCuboid(67, 94, 50)
	shapesSlice[4] = Circle(19.45)
	//shapesSlice = append(shapesSlice, 213.123)

	for _, v := range shapesSlice {
		// s, ok := v.(Shaper)
		// if ok {
		// 	//PrintShape(v.(Shaper))
		// 	PrintShape(s)
		// } else {
		// 	println("invalid independent object")
		// }

		// switch s := v.(type) {
		// case Shaper:
		// 	PrintShape(s)
		// default:
		// 	println("invalid independent object")
		// }
		PrintShape(v)
		println()
	}

}

func PrintShape(shaper Shaper) {
	fmt.Printf("Area of %s: %0.2f  \n", shaper.What(), shaper.Area())
	fmt.Printf("Perimeter of %s: %0.2f \n", shaper.What(), shaper.Perimeter())
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
