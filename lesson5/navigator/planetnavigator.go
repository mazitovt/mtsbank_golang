package navigator

import (
	"fmt"
	"mtsbank_golang/lesson5/distance/path"
	"mtsbank_golang/lesson5/distance/points"
	"mtsbank_golang/lesson5/geocoding"
)

type PlanetNavigator struct {
	geocoder geocoding.Geocoder
	path     path.PathOnSphere
	location points.PointOnSphere
	position int
	passed   float64
	distance float64
}

func NewPlanetNavigator(geocoder geocoding.Geocoder, path path.PathOnSphere) (*PlanetNavigator, error) {
	if d, e := path.Distance(); e != nil {
		return nil, e
	} else {
		return &PlanetNavigator{geocoder: geocoder, path: path, position: -1, passed: 0, distance: d}, nil
	}
}

func (p *PlanetNavigator) DistancePassed() float64 {
	return p.passed
}

func (p *PlanetNavigator) DistanceLeft() float64 {
	return p.distance - p.passed
}

func (p *PlanetNavigator) CurrentLocation() (string, error) {
	return p.path.PointAt(p.position)
}

func (p *PlanetNavigator) CurrentAddress() (string, error) {
	pointOnSph, e := p.path.PointOnSphereAt(p.position)
	fmt.Println(e)
	data, e := p.geocoder.ReverseGeocode(pointOnSph)
	fmt.Println(e)
	return data.Address, nil
}

func (p *PlanetNavigator) MoveNext() bool {
	d, _ := p.NextDistance()
	p.position++
	p.passed += d
	return p.position < p.path.CountPoints()
}

func (p *PlanetNavigator) NextDistance() (float64, error) {
	return p.path.DistanceBetween(p.position, p.position+1)
}

func (p *PlanetNavigator) ResetRoute() {
	p.position = -1
	p.passed = 0
}

func (p *PlanetNavigator) getGeocodeData() (geocoding.GeocodeData, error) {
	d, _ := p.geocoder.ReverseGeocode(p.location)
	return d, nil
}
