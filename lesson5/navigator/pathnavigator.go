package navigator

import "mtsbank_golang/lesson5/distance/path"

type PathNavigator struct {
	path     path.Path
	position int
	passed   float64
	distance float64
}

func (n *PathNavigator) DistancePassed() float64 {
	return n.passed
}

func (n *PathNavigator) DistanceLeft() float64 {
	return n.distance - n.passed
}

func (n *PathNavigator) CurrentLocation() (string, error) {
	return n.path.PointAt(n.position)
}

func (n *PathNavigator) NextLocation() bool {
	n.position++
	t, _ := n.path.DistanceBetween(n.position, n.position-1)
	n.passed += t
	return n.position < n.path.CountPoints()
}

func (n *PathNavigator) ResetRoute() {
	n.position = -1
	n.passed = 0
	n.distance, _ = n.path.Distance()
}

func NewAnyPathNavigator(path path.Path) (a *PathNavigator, err error) {
	if d, e := path.Distance(); err != nil {
		return a, e
	} else {
		return &PathNavigator{path, -1, 0, d}, err
	}
}
