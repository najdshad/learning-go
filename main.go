package main

import (
	"fmt"

	"learning-go/src"
)

func main() {
	// src.ParseArgs()

	myCircle := src.CalcArea(src.Circle, src.ShapeParameters{Radius: 5.3})
	fmt.Printf("area of myShape is: %f", myCircle)
}
