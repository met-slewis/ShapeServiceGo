package ShapeServiceGo

import "github.com/paulmach/orb"

const (
	subBucketName     = "weather-event-sub"
	locationsFilename = "locations.json"
)

// PointType lat/long
type PointType [2]float64
type WarningPolyType []PointType
type WarningPolysType []WarningPolyType

func (p *PointType) GetLon() float64 {
	return p[1]
}
func (p *PointType) GetLat() float64 {
	return p[0]
}
func (p *PointType) SetLon(lon float64) {
	p[1] = lon
}
func (p *PointType) SetLat(lat float64) {
	p[0] = lat
}
func (p *PointType) SetLatLon(lat, lon float64) {
	p[0] = lat
	p[1] = lon
}

func (p *PointType) ToGeoJsonPoint() orb.Point {
	point := orb.Point{
		p[1],
		p[0],
	}
	return point
}

// ToGeoJsonMultiPoly  convert each of the poly's in the warning to
// a geoJson polygon.  Add each geoJson poly to a geoJson MultiPolygon
// lat/lon => lon/lat
// each poly from the warning becomes the initial poly (ring) in a new geoJson poly
func (p *WarningPolysType) ToGeoJsonMultiPoly() orb.MultiPolygon {
	outPolys := orb.MultiPolygon{}
	// a warning polygon becomes the first 'ring' in a geoJson polygon
	for _, inPoly := range *p {
		outPoly := inPoly.ToGeoJsonPoly()
		outPolys = append(outPolys, outPoly)
	}
	return outPolys
}

func (p *WarningPolyType) ToGeoJsonPoly() orb.Polygon {
	orbRing := orb.Ring{}
	for _, inPoint := range *p {
		orbRing = append(orbRing, inPoint.ToGeoJsonPoint())
	}
	outPoly := orb.Polygon{}
	outPoly = append(outPoly, orbRing)
	return outPoly
}

func (p *WarningPolysType) CastToWarningPolysType(in [][][2]float64) WarningPolysType {
	for _, poly := range in {
		var wPoly WarningPolyType
		for _, point := range poly {
			wPoly = append(wPoly, point)
		}
		*p = append(*p, wPoly)
	}
	return *p
}
