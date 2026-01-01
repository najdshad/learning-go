package src

import "math"

// Shape : define shape options as enum type
type Shape int

const (
	Circle Shape = iota
	Triangle
	Rectangle
)

// ShapeParameters struct holds optional argument values
type ShapeParameters struct {
	Radius float64
	Base   float64
	Height float64
	Width  float64
}

func CalcArea(s Shape, p ShapeParameters) float64 {
	switch s {
	case Circle:
		return math.Pi * p.Radius * p.Radius
	case Triangle:
		return p.Base * p.Height / 2.0
	case Rectangle:
		return p.Height * p.Width
	default:
		return 0
	}
}
