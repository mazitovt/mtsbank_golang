package geocoding

import "mtsbank_golang/tests_mocks/distance/points"

type Geocoder interface {
	Geocode(address string) (point points.PointOnSphere, err error)
	ReverseGeocode(point points.PointOnSphere) (data GeocodeData, err error)
}

type GeocodeData struct {
	Point      points.PointOnSphere
	Address    string
	PostalCode string
}
