package geocoding

import "mtsbank_golang/lesson5/distance/points"

type Geocoder interface {
	Geocode(address string) (point points.PointOnSphere, err error)
	ReverseGeocode(point points.PointOnSphere) (data GeocodeData, err error)
}

type GeocodeData struct {
	Point      points.PointOnSphere
	Address    string
	PostalCode string
}
