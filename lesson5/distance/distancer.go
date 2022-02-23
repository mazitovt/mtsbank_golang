package distance

type Distancer interface {
	Distance() (float64, error)
}
