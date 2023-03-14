package ShapeServiceGo

import "github.com/ctessum/geom"

func PolysOverlap(warnPoly1 WarningPolyType, warnPoly2 WarningPolyType) bool {

	poly1 := WarnToGeom(warnPoly1)
	poly2 := WarnToGeom(warnPoly2)

	intersect := poly1.Intersection(poly2)

	if intersect.Area() > 0.0 {
		return true
	}

	return false
}

func WarnToGeom(warnPoly WarningPolyType) geom.Polygonal {
	poly := geom.Polygon{}

	// type Polygon []Path
	// type Path []Point
	// type Point struct {
	//  X, Y float64
	//}

	return poly
}
