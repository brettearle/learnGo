package structs

import "math"

// interface at the top
// make functions that take a interface then it
// will be generic aslong as whats passed as a arg
// implicitly implements that interface
// so shape is a parameter in area tests and then
// it allowes all the concrete shapes (rect, circle, tri)
// to be passed in because they all implement Shape interface
type Shape interface {
	Area() float64
}

// this is the struct
type Rectangle struct {
	Width  float64
	Height float64
}

// this is a method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// this is the struct
type Circle struct {
	Radius float64
}

// this is the method
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
