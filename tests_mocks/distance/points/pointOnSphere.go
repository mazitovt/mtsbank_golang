package points

import (
	"encoding/json"
)

type PointOnSphere struct {
	lat float64
	lon float64
}

func (p PointOnSphere) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(
		struct {
			X         float64 `json:"lat"`
			Y         float64 `json:"lon"`
			CreatedAt string  `json:"createdAt,omitempty"`
		}{p.lat, p.lon, ""})

	if err != nil {
		return nil, err
	}

	return j, err
}

func (p PointOnSphere) Lat() float64 {
	return p.lat
}

func (p PointOnSphere) Lon() float64 {
	return p.lon
}

func NewPointOnSphereFromDegrees(lat, lon float64) (p *PointOnSphere, err error) {

	if lat > 90 || lat < -90 {
		err = &InvalidLatitudeError{}
		return
	}

	if lon > 180 || lon < -180 {
		return nil, &InvalidLongitudeError{}
	}

	p = &PointOnSphere{lat, lon}

	return
}
