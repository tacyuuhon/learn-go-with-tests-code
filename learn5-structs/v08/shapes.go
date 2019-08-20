package structs

import "math"

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method
// Calc Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method
// Calc Circle
func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area method
// Calc Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Shape interface
type Shape interface {
	Area() float64
}

// Premeter func
func Premeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area func
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
