package navigator

import (
	"errors"
	"mtsbank_golang/tests_mocks/distance"
	"mtsbank_golang/tests_mocks/distance/points"
	"mtsbank_golang/tests_mocks/geocoding"
)

type PathInfo struct {
	placeStart  geocoding.GeocodeData
	placeFinish geocoding.GeocodeData
}

func (p PathInfo) PlaceStart() geocoding.GeocodeData {
	return p.placeStart
}

func (p PathInfo) PlaceFinish() geocoding.GeocodeData {
	return p.placeFinish
}

type MyNavigator struct {
	distances []distance.Distancer
	geocoder  geocoding.Geocoder
}

//func NewNavigator(distances []Distancer) *MyNavigator {
//	return &MyNavigator{distances: distances}
//}

func NewNavigator(geocoding geocoding.Geocoder) *MyNavigator {
	return &MyNavigator{geocoder: geocoding}
}

//func (n MyNavigator) Path() (path float64, err error) {
//	for _, dist := range n.distances {
//		pathLocal, err1 := dist.Distancer()
//
//		if err1 != nil {
//			return path, err1
//		}
//		path += pathLocal
//	}
//
//	return
//
//}

func (n MyNavigator) PathInfo(point1 points.PointOnSphere, point2 points.PointOnSphere) (info PathInfo, err error) {
	data1, err := n.geocoder.ReverseGeocode(point1)
	if err != nil {
		err = errors.New("Error geocode first point")
	}

	data2, err := n.geocoder.ReverseGeocode(point2)
	if err != nil {
		err = errors.New("Error geocode second point")
	}

	info = PathInfo{placeStart: data1, placeFinish: data2}
	return
}
