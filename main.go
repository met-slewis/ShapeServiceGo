package main

import (
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/planar"
)

func main() {
	ExampleArea()
	ExamplePointInPoly()
}

func ExampleArea() {
	poly := orb.Polygon{
		{
			{-122.4163816, 37.7792782},
			{-122.4162786, 37.7787626},
			{-122.4151027, 37.7789118},
			{-122.4152143, 37.7794274},
			{-122.4163816, 37.7792782},
		},
	}
	a := geo.Area(poly)

	fmt.Printf("%f m^2\n", a)
	// Output:
	// 6073.368008 m^2
}

func ExamplePointInPoly() {
	var point orb.Point
	point = [2]float64{-46.2, 170.1}
	poly := Poly1()
	ok := isPointInPoly(point, poly)

	fmt.Printf("Point in Poly: %t\n", ok)
}

func isPointInPoly(point orb.Point, poly orb.Polygon) bool {

	ring := poly[0]

	contains := planar.RingContains(ring, point)

	return contains
}

func WarningPoly() orb.Polygon {
	return Poly1()
}
