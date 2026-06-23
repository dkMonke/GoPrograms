// main.go — Day 10: Interfaces.
// Shape is an interface requiring an Area() float64 method. Rectangle and Circle
// both satisfy Shape implicitly. totalArea accepts a []Shape, demonstrating
// polymorphism — any type with Area() can be passed without explicit "implements".
package main

import "fmt"

// Shape is the abstraction for any geometric figure that can report its area.
// Any type implementing Area() float64 satisfies Shape implicitly, enabling
// polymorphic handling by functions such as totalArea.
type Shape interface {
	Area() float64
}

// Rectangle is a Shape defined by its Width and Height.
type Rectangle struct {
	Width, Height float64
}

// Area implements the Shape interface for Rectangle, returning the product of
// its Width and Height.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a Shape defined by its Radius.
type Circle struct {
	Radius float64
}

// Area implements the Shape interface for Circle, returning pi * r^2 using an
// approximation of pi (3.14159).
func (c Circle) Area() float64 {

	return 3.14159 * c.Radius * c.Radius

}

// totalArea sums the Area of every Shape in the shapes slice and returns the
// combined total. It works with any mix of concrete Shape implementations.
func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

// main constructs a slice containing a Rectangle and a Circle and prints their
// combined area via totalArea, demonstrating interface-based polymorphism.
func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
	}
	fmt.Printf("total : %.2f\n", totalArea(shapes))
}
