package navigator

type Navigator interface {
	DistancePassed() float64
	DistanceLeft() float64
	CurrentLocation() (string, error)
	MoveNext() bool
	NextDistance() (float64, error)
	ResetRoute()
	CurrentAddress() (string, error)
}
