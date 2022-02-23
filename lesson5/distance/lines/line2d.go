package lines

import (
	"encoding/json"
	"mtsbank_golang/lesson5/distance/points"
	"mtsbank_golang/lesson5/distance/utils"
)

type Line2d struct {
	start points.Point2d
	end   points.Point2d
}

func (l Line2d) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(
		struct {
			Point1 points.Point2d `json:"p1"`
			Point2 points.Point2d `json:"p2"`
		}{l.start, l.end})

	if err != nil {
		return nil, err
	}

	return j, err
}

func (l Line2d) Point1() points.Point2d {
	return l.start
}

func (l Line2d) Point2() points.Point2d {
	return l.end
}

func NewLine2d(point1 points.Point2d, point2 points.Point2d) *Line2d {
	return &Line2d{start: point1, end: point2}
}

func (l Line2d) Distance() float64 {
	return utils.ShortestDistanceTwoPoints2D(l.start.X(), l.start.Y(), l.end.X(), l.end.Y())
}
