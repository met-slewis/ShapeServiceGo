package ShapeServiceGo

import (
	"encoding/json"
	"fmt"
	"github.com/MetServiceDev/WeatherEventLib"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/planar"
	"os"
	"testing"
)

func ExampleArea() {
	poly := orb.Polygon{
		{
			{122.4163816, -37.7792782},
			{122.4162786, -37.7787626},
			{122.4151027, -37.7789118},
			{122.4152143, -37.7794274},
			{122.4163816, -37.7792782},
		},
	}
	a := geo.Area(poly)

	fmt.Printf("%f m^2\n", a)
	// Output:
	// 6073.368008 m^2
}

func TestIsPointInPoly(t *testing.T) {
	point := PointType{-46.2, 170.1}
	poly := Poly2()
	ok := IsPointInPoly(point, poly)
	fmt.Printf("Point in Poly: %t\n", ok)
}

func TestIsPointInPolyOrb(t *testing.T) {
	point := orb.Point{170.1, -46.2}
	poly := Poly1Orb()
	multiPoly := orb.MultiPolygon{}
	multiPoly = append(multiPoly, poly)

	ok := IsPointInPolyOrb(point, multiPoly)

	fmt.Printf("Point in Poly: %t\n", ok)
}

func TestPointInMultiPoly(t *testing.T) {
	warning := getAuckThreePolyWarning()
	var warningPolys WarningPolysType
	warningPolys.CastToWarningPolysType(warning.Polygons)
	pointsIn := getPointsInsideWarning()

	for _, point := range pointsIn {
		if !IsPointInPoly(point, warningPolys) {
			t.Errorf("Point %v is incorrectly seen as outside of polys", point)
		}
	}
}

func TestPointOutOfMultiPoly(t *testing.T) {
	warning := getAuckThreePolyWarning()
	var warningPolys WarningPolysType
	warningPolys.CastToWarningPolysType(warning.Polygons)
	pointsIn := getPointsOutsideWarning()

	for _, point := range pointsIn {
		if IsPointInPoly(point, warningPolys) {
			t.Errorf("Point %v is incorrectly seen as in poly", point)
		}
	}
}

func isPointInPoly(point orb.Point, poly orb.Polygon) bool {
	return planar.PolygonContains(poly, point)
}

func WarningPoly() orb.Polygon {
	return Poly1Orb()
}

func getAuckThreePolyWarning() WeatherEventLib.LandWarning {
	filename := "./res/Auck-Wind-Red.json"
	jsonBytes, _ := os.ReadFile(filename)
	var warning WeatherEventLib.LandWarning
	json.Unmarshal(jsonBytes, &warning)
	return warning
}

func getPointsInsideWarning() []PointType {
	points := []PointType{
		{-36.18742135683066, 175.40863696763188},
		{-36.200117857925896, 175.07375979139897},
		{-36.14387486954988, 174.60178524771243},
		{-36.48253371989178, 174.21521562145438},
		{-37.32540250461747, 174.71865513471886},
		{-37.13213489287093, 175.2490646219112},
		{-36.79812558146074, 175.10747225880476},
		{-36.567419632980105, 174.6647151868712},
		{-36.36136520524066, 174.7636050912609},
	}
	return points
}

func getPointsOutsideWarning() []PointType {
	points := []PointType{
		{-36.16832790839182, 174.30176928915148},
		{-35.76353909471764, 174.04521107374921},
		{-36.52409036487239, 175.38176776844227},
		{-37.36026898631682, 175.83613215741053},
		{-37.63537311379954, 174.89153250666038},
		{-38.597947715899444, 175.9262412978665},
	}
	return points
}
