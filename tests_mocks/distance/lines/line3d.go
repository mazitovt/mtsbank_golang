package lines

import (
	"mtsbank_golang/tests_mocks/distance/points"
	"mtsbank_golang/tests_mocks/distance/utils"
)

type Line3d struct {
	start points.Point3d
	end   points.Point3d
}

func (l Line3d) Point1() points.Point3d {
	return l.start
}

func (l Line3d) Point2() points.Point3d {
	return l.end
}

func NewLine3d(point1 points.Point3d, point2 points.Point3d) *Line3d {
	return &Line3d{start: point1, end: point2}
}

func (l Line3d) Distance() (distance float64) {
	return utils.ShortestDistanceTwoPoints3D(l.start.X(), l.start.Y(), l.start.Z(), l.end.X(), l.end.Y(), l.end.Z())
}
