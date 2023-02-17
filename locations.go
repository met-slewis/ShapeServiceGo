package ShapeServiceGo

import lib "github.com/MetServiceDev/WeatherEventLib"

func GetAffectedLocations(allLocations lib.LocationsType, polys WarningPolysType) lib.LocationsType {

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
