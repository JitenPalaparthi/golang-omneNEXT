package main

import (
	"omnenext-shapes/circle"
	"omnenext-shapes/cuboid"
	"omnenext-shapes/demo"
	"omnenext-shapes/rect"
	"omnenext-shapes/shapes"
	"omnenext-shapes/square"
)

func main() {

	shapes.Greet()

	demo.Greet()

	t1 := demo.T1{A: 100, C: 300}
	t1.SetB(200)
	t1.SetD(400)
	t1.Print()

	shapesSlice := make([]shapes.Shaper, 5)
	shapesSlice[0] = rect.NewRect(10.23, 7.89)
	shapesSlice[1] = square.Square(32.23)
	shapesSlice[2] = circle.Circle(56.7)
	shapesSlice[3] = cuboid.NewCuboid(67, 94, 50)
	shapesSlice[4] = circle.Circle(19.45)
	//shapesSlice = append(shapesSlice, 213.123)

	for _, v := range shapesSlice {
		shapes.PrintShape(v)
		println()
	}

	//shapes.greet()

}
