package shapes

import "math"

// In Go, interface resolution is implicit
// if the type passed in matches with what the interface is asking for, it'll compile
type Shape interface {
	Area() float64
}

// struct: named collection of fields where you can store data
type Rectangle struct {
	Width  float64
	Height float64
}

// when method is called on a variable of that type, you'll get the reference to
// its data via the receiverName variable, here it is "r" (this)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

// func Area(circle Circle) float64 {
// 	return 0
// }
