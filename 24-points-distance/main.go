// Task 24
// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) DistanceTo(another *Point) float64 {
	return math.Sqrt(math.Pow(another.x-p.x, 2) + math.Pow(another.y-p.y, 2))
}

func main() {
	p1 := NewPoint(8, -2)
	p2 := NewPoint(-2, 8)

	fmt.Println(p1.DistanceTo(p2))
	fmt.Println(p2.DistanceTo(p1))
}
