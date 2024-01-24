package main

import "fmt"

type GeometryFigure interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radio float64
}

func (c Circle) Area() (result float64) {
	result = 3.14 * c.Radio * c.Radio

	return
}

func (c Circle) Perimeter() (result float64) {
	result = 2.0 * 3.14 * c.Radio

	return
}

type Square struct {
	side float64
}

func (s Square) Area() (result float64) {
	result = s.side * s.side

	return
}

func (s Square) Perimeter() (result float64) {
	result = s.side * 4.0

	return
}

func PrintAreaAndPerimeterGeometryFigure(gf GeometryFigure) {
	fmt.Printf("Area: %f\n", gf.Area())
	fmt.Printf("Perimeter: %f\n", gf.Perimeter())
}
func main() {

	dataCir := Circle{
		Radio: 1.0,
	}
	PrintAreaAndPerimeterGeometryFigure(dataCir)

	dataSqu := Square{
		side: 1.0,
	}
	PrintAreaAndPerimeterGeometryFigure(dataSqu)
}
