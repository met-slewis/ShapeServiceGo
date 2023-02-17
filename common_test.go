package ShapeServiceGo

import (
	"fmt"
	"testing"
)

func TestWarningPolyType_ToGeoJsonPoly(t *testing.T) {
	poly := createWarningPolygon()
	geoJsonPoly := poly.ToGeoJsonPoly()
	if len(geoJsonPoly) != 1 {
		t.Errorf("should be exactly 1 element (ring) in poly")
	}
}

func TestWarningPolysType_ToGeoJsonMultiPoly(t *testing.T) {
	polys := createWarningPolygons()
	geoMultiPoly := polys.ToGeoJsonMultiPoly()
	if len(geoMultiPoly) != 3 {
		t.Errorf("geoJson polys needs 3 polys")
	}
	fmt.Printf("%v", geoMultiPoly)
}

func createWarningPolygon() WarningPolyType {
	return WarningPolyType{{-42.0, 174.0}, {-42.1, 174.1}, {-42.2, 174.2}, {-42.3, 174.3}, {-42.0, 174.0}}
}

func createWarningPolygons() WarningPolysType {
	ring1 := []PointType{{-42.0, 174.0}, {-42.1, 174.1}, {-42.2, 174.2}, {-42.3, 174.3}, {-42.0, 174.0}}
	ring2 := []PointType{{-43.0, 174.0}, {-43.1, 174.1}, {-43.2, 174.2}, {-43.3, 174.3}, {-43.0, 174.0}}
	ring3 := []PointType{{-44.0, 174.0}, {-44.1, 174.1}, {-44.2, 174.2}, {-44.3, 174.3}, {-44.0, 174.0}}

	poly := WarningPolysType{}
	poly = append(poly, ring1)
	poly = append(poly, ring2)
	poly = append(poly, ring3)
	return poly
}
