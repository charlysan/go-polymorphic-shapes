package main

import (
	"fmt"
	"math"
	"reflect"
)

// Shaper defines method signatures for shapes
type Shaper interface {
	Area() float32
	Perimeter() float32
	Dimensions() string
}

type Rectangle struct {
	length float32
	width  float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Dimensions() string {
	return fmt.Sprintf("length %.2f cm and width %.2f cm", r.length, r.width)
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius

}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Dimensions() string {
	return fmt.Sprintf("radius %.2f cm", c.radius)
}

func calculateArea(s Shaper) float32 {
	return s.Area()
}

func calculatePerimeter(s Shaper) float32 {
	return s.Perimeter()
}

func displayDimensions(s Shaper) string {
	return s.Dimensions()
}

func main() {
	s := Rectangle{2, 3}
	c := Circle{1}

	shapesCollection := []Shaper{s, c}

	for _, shape := range shapesCollection {
		fmt.Printf(
			"A %s of %s has an area of %.2f cm2 and a perimeter of %.2f cm.\n",
			reflect.TypeOf(shape).Name(),
			displayDimensions(shape),
			calculateArea(shape),
			calculatePerimeter(shape),
		)
	}
}
