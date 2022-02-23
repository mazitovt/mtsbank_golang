package navigator

import "mtsbank_golang/lesson5/geocoding"

type PlanetNavigator interface {
	GetGeocodeData() geocoding.GeocodeData
}
