package path

type Path interface {
	Distance() (float64, error)
	DistanceBetween(int, int) (float64, error)
	CountPoints() int
	PointAt(int) (string, error)
}
