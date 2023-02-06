package ShapeServiceGo

import (
	"encoding/json"
	lib "github.com/MetServiceDev/WeatherEventLib"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

const (
	resDir        = "./res/"
	locationsFile = resDir + "locations.json"
)

func TestGetAffectedLocations(t *testing.T) {
	inputLocations := getInputLocations()
	poly := getTestPoly()

	affectedLocations := GetAffectedLocations(inputLocations, poly)

	numAffectedLocations := len(affectedLocations.Locations)
	if numAffectedLocations != 3 {
		t.Errorf("Should be exactly 3 affected location, there were %d", numAffectedLocations)
	}

}

func getTestPoly() polyType {
	return Poly2()
}

type TmpLocationsType struct {
	Locations []lib.LocationType `json:"locations"`
}

func getInputLocations() lib.LocationsType {
	locationsData, err := os.ReadFile(locationsFile)
	if err != nil {
		logrus.Errorf("Error reading locations data. %s", err)
	}

	var locations TmpLocationsType
	err = json.Unmarshal(locationsData, &locations)
	if err != nil {
		logrus.Errorf("Unable to unmarshall into LocationsType. %s", err.Error())
	}

	locationsMap := lib.LocationsType{
		Locations: make(map[string]lib.LocationType),
	}

	for _, loc := range locations.Locations {
		locationsMap.Locations[loc.LocationId] = loc
	}
	return locationsMap
}
