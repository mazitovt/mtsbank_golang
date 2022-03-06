package lines

import (
	"mtsbank_golang/tests_mocks/distance/points"
	"mtsbank_golang/tests_mocks/distance/utils"
)

type LineOnSphere struct {
	start, end points.PointOnSphere
}

func (l *LineOnSphere) Start() points.PointOnSphere {
	return l.start
}

func (l *LineOnSphere) End() points.PointOnSphere {
	return l.end
}

func NewLineOnSphere(start, end points.PointOnSphere) *LineOnSphere {
	return &LineOnSphere{start, end}
}

func (l LineOnSphere) Distance(r float64) (distance float64) {

	lat1 := utils.DegreesToRadians(l.start.Lat())
	lon1 := utils.DegreesToRadians(l.start.Lon())
	lat2 := utils.DegreesToRadians(l.end.Lat())
	lon2 := utils.DegreesToRadians(l.end.Lon())

	return utils.ShortestDistanceTwoPointsOnSphere(r, lat1, lon1, lat2, lon2)
}
