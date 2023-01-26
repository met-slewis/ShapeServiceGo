package ShapeServiceGo

import (
	lib "github.com/MetServiceDev/WeatherEventLib"
	//  ums "github.com/met-slewis/WeatherUMS"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

//func PointsInPoly(locations lib.LocationsType, polygon [][][2]float64) lib.LocationsType {
//  affectedLocations := lib.LocationsType{
//    Locations: make(map[string]lib.LocationType, 0),
//  }
//
//  for _, loc := range locations.Locations {
//
//  }
//
//}

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

type ePoly [][][][2]float64

func GetAffectedLocations(allLocations lib.LocationsType, poly [][][2]float64) lib.LocationsType {
	polys := make([][][][2]float64, 0)
	polys = append(polys, poly)

	affectedLocations := lib.LocationsType{
		Locations: make(map[string]lib.LocationType, 0),
	}

	for _, loc := range allLocations.Locations {
		if IsPointInPoly(loc.Point(), polys) {
			affectedLocations.Locations[loc.LocationId] = loc
		}
	}
	return affectedLocations
}
