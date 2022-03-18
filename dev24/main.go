// 24
// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде
// структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Create Point struct.
type Point struct {
	x, y float64 // Fields for x, y values
}

// NewPoint — creates a new Point and returns a pointer to it (aka constructor).
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo — calculates distance to another point.
func (p *Point) DistanceTo(another *Point) float64 {
	return math.Sqrt(math.Pow(another.x-p.x, 2) + math.Pow(another.y-p.y, 2))
}

func main() {
	// Create two points.
	p1 := NewPoint(8, -2)
	p2 := NewPoint(-2, 8)

	// Calculate and print distances from p1 to p2 and from p2 to p1. Distances must be the same.
	fmt.Println(p1.DistanceTo(p2))
	fmt.Println(p2.DistanceTo(p1))
}
