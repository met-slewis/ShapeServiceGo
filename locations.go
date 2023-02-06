package ShapeServiceGo

import lib "github.com/MetServiceDev/WeatherEventLib"

type polyType [][][2]float64
type ePolyType [][][][2]float64

func GetAffectedLocations(allLocations lib.LocationsType, poly [][][2]float64) lib.LocationsType {
	polys := make(ePolyType, 0)
	polys = append(polys, poly)

	affectedLocations := lib.LocationsType{
		Locations: make(map[string]lib.LocationType, 0),
	}

	for _, loc := range allLocations.Locations {
		if loc.AllLocations() {
			affectedLocations.Locations[loc.LocationId] = loc
		} else if IsPointInPoly(loc.Point(), polys) {
			affectedLocations.Locations[loc.LocationId] = loc
		}
	}
	return affectedLocations
}
