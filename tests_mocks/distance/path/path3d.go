package path

import (
	"errors"
	"fmt"
	"mtsbank_golang/tests_mocks/distance/lines"
	"mtsbank_golang/tests_mocks/distance/points"
)

type Path3d struct {
	points []points.Point3d
}

func NewPath3d(points []points.Point3d) *Path3d {
	return &Path3d{points}
}

func (p *Path3d) AddPoint(point points.Point3d) *Path3d {
	p.points = append(p.points, point)
	return p
}

func (p *Path3d) CountPoints() int {
	return len(p.points)
}

func (p *Path3d) PointAt(i int) (s string, err error) {

	if i < 0 || i >= len(p.points) {
		err = errors.New("invalid index error")
	} else {
		s = fmt.Sprintf("%v", p.points[i])
	}

	return
}

func (p *Path3d) Distance() (distance float64, err error) {

	l := len(p.points)

	if l < 2 {
		return distance, &InvalidNumberOfPointForDistance{l}
	}

	for i := 0; i < l-1; i++ {
		distance += p.innerDistanceBetween(p.points[i], p.points[i+1])
	}

	return
}

func (p *Path3d) innerDistanceBetween(start, end points.Point3d) float64 {
	return lines.NewLine3d(start, end).Distance()
}

func (p *Path3d) DistanceBetween(start, end int) (d float64, err error) {

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
