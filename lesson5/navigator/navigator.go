package navigator

type Navigator interface {
	DistancePassed() float64
	DistanceLeft() float64
	CurrentLocation() (string, error)
	NextLocation() bool
	ResetRoute()
}
