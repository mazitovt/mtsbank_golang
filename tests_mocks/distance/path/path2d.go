package path

import (
	"errors"
	"fmt"
	"mtsbank_golang/tests_mocks/distance/lines"
	"mtsbank_golang/tests_mocks/distance/points"
)

type Path2d struct {
	points []points.Point2d
}

func NewPath2d(points []points.Point2d) *Path2d {
	return &Path2d{points}
}

func (p *Path2d) AddPoint(point points.Point2d) *Path2d {
	p.points = append(p.points, point)
	return p
}

func (p *Path2d) CountPoints() int {
	return len(p.points)
}

func (p *Path2d) PointAt(i int) (s string, err error) {

	if i < 0 || i >= len(p.points) {
		err = errors.New("invalid index error")
	} else {
		s = fmt.Sprintf("%v", p.points[i])
	}

	return
}

func (p *Path2d) Distance() (distance float64, err error) {

	l := len(p.points)

	if l < 2 {
		return distance, &InvalidNumberOfPointForDistance{l}
	}

	for i := 0; i < l-1; i++ {
		distance += p.innerDistanceBetween(p.points[i], p.points[i+1])
	}

	return
}

func (p *Path2d) innerDistanceBetween(start, end points.Point2d) float64 {
	return lines.NewLine2d(start, end).Distance()
}

func (p *Path2d) DistanceBetween(start, end int) (d float64, err error) {

	if start > end {
		start, end = end, start
	}
	if start >= 0 && start < len(p.points)-1 && start == end-1 {
		d = p.innerDistanceBetween(p.points[start], p.points[end])
	} else {
		err = errors.New("invalid indexes error")
	}

	return
}
