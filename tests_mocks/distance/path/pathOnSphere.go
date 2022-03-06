package path

import (
	"errors"
	"fmt"
	"mtsbank_golang/tests_mocks/distance/lines"
	"mtsbank_golang/tests_mocks/distance/points"
)

type PathOnSphere struct {
	radius float64
	points []points.PointOnSphere
}

func NewPathOnSphere(r float64, points []points.PointOnSphere) *PathOnSphere {
	return &PathOnSphere{r, points}
}

func (p *PathOnSphere) AddPoint(point points.PointOnSphere) *PathOnSphere {
	p.points = append(p.points, point)
	return p
}

func (p *PathOnSphere) CountPoints() int {
	return len(p.points)
}

func (p *PathOnSphere) PointAt(i int) (s string, err error) {

	if i < 0 || i >= len(p.points) {
		err = errors.New("invalid index error")
	} else {
		s = fmt.Sprintf("%v", p.points[i])
	}

	return
}

func (p *PathOnSphere) PointOnSphereAt(i int) (s points.PointOnSphere, err error) {

	if i < 0 || i >= len(p.points) {
		err = errors.New("invalid index error")
	} else {
		s = p.points[i]
	}

	return
}

func (p *PathOnSphere) Distance() (distance float64, err error) {

	l := len(p.points)

	if l < 2 {
		return distance, &InvalidNumberOfPointForDistance{l}
	}

	for i := 0; i < l-1; i++ {
		distance += p.innerDistanceBetween(p.points[i], p.points[i+1])
	}

	return
}

func (p *PathOnSphere) innerDistanceBetween(start, end points.PointOnSphere) float64 {
	return lines.NewLineOnSphere(start, end).Distance(p.radius)
}

func (p *PathOnSphere) DistanceBetween(start, end int) (d float64, err error) {

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
