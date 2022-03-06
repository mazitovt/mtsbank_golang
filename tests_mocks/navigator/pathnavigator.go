package navigator

import "mtsbank_golang/tests_mocks/distance/path"

type PathNavigator struct {
	path     path.Path
	position int
	passed   float64
	distance float64
}

func (p *PathNavigator) CurrentAddress() (string, error) {
	return p.CurrentLocation()
}

func (p *PathNavigator) DistancePassed() float64 {
	return p.passed
}

func (p *PathNavigator) DistanceLeft() float64 {
	return p.distance - p.passed
}

func (p *PathNavigator) CurrentLocation() (string, error) {
	return p.path.PointAt(p.position)
}

func (p *PathNavigator) MoveNext() bool {
	d, _ := p.NextDistance()
	p.position++
	p.passed += d
	return p.position < p.path.CountPoints()
}

func (p *PathNavigator) NextDistance() (float64, error) {
	return p.path.DistanceBetween(p.position, p.position+1)
}

func (p *PathNavigator) ResetRoute() {
	p.position = -1
	p.passed = 0
}

func NewPathNavigator(path path.Path) (a *PathNavigator, err error) {
	if d, e := path.Distance(); err != nil {
		return a, e
	} else {
		return &PathNavigator{path, -1, 0, d}, err
	}
}
