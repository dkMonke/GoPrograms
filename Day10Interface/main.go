// main.go — Day 10: Interfaces.
// Shape is an interface requiring an Area() float64 method. Rectangle and Circle
// both satisfy Shape implicitly. totalArea accepts a []Shape, demonstrating
// polymorphism — any type with Area() can be passed without explicit "implements".
package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	return 3.14159 * c.Radius * c.Radius

}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
func main() {
	shapes := []Shape{
		Rectangle{Width: 3, Height: 4},
		Circle{Radius: 5},
	}
	fmt.Printf("total : %.2f\n", totalArea(shapes))
}
