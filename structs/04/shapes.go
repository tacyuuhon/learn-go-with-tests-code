package structs

// Rectangle struct
type Rectangle struct {
	width  float64
	height float64
}

// Premeter func
func Premeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Area func
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
