package ShapeServiceGo

import (
	//  ums "github.com/met-slewis/WeatherUMS"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

func IsPointInPoly(point PointType, polygons WarningPolysType) bool {
	p := point.ToGeoJsonPoint()
	multiPoly := polygons.ToGeoJsonMultiPoly()
	return IsPointInPolyOrb(p, multiPoly)
}

func IsPointInPolyOrb(point orb.Point, multiPoly orb.MultiPolygon) bool {
	return planar.MultiPolygonContains(multiPoly, point)
}
