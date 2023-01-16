package main

import (
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
)

func main() {
	ExampleArea()
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

	fmt.Printf("%f m^2", a)
	// Output:
	// 6073.368008 m^2
}

func ExamplePointInPoly() {
	point := orb.Point{}

}

func WarningPoly() orb.Polygon {
	return Poly1()
}
