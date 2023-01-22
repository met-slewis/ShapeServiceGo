package main

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

func IsPointInPoly(point [2]float64, polygons [][][][2]float64) bool {
	var p orb.Point
	p = point

	multiPoly := make([]orb.Polygon, 0)
	for _, poly := range polygons {
		oPoly := ToPoly(poly)
		multiPoly = append(multiPoly, oPoly)
	}

	return IsPointInPolyOrb(p, multiPoly)
}

// ToPoly  A polygon is defined as an array of ring, a Ring is a
// type of linestring that sits on the earth
func ToPoly(in [][][2]float64) orb.Polygon {
	oPoly := orb.Polygon{}
	for _, ring := range in {
		// for each Ring (the first is the poly,
		// subsequent rings are the holes)
		oRing := orb.Ring{}
		for _, p := range ring {
			oRing = append(oRing, p)
		}
		oPoly = append(oPoly, oRing)
	}
	return oPoly
}

func IsPointInPolyOrb(point orb.Point, multiPoly orb.MultiPolygon) bool {
	contains := planar.MultiPolygonContains(multiPoly, point)
	return contains
}
