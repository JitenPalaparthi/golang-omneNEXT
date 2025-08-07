package main

import (
	"github.com/JitenPalaparthi/omnenext-shapes/circle"
	"github.com/JitenPalaparthi/omnenext-shapes/cuboid"
	"github.com/JitenPalaparthi/omnenext-shapes/demo"
	"github.com/JitenPalaparthi/omnenext-shapes/rect"
	"github.com/JitenPalaparthi/omnenext-shapes/shapes"
	"github.com/JitenPalaparthi/omnenext-shapes/square"
	_ "github.com/gin-gonic/gin"
)

func main() {

	shapes.Greet()

	demo.Greet()

	demo.Sayhi()

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
