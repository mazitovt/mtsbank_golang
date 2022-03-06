package car

import "fmt"

type NegativeFuelValueError struct {
	value float64
}

func (e *NegativeFuelValueError) Error() string {
	return fmt.Sprintf("zero or negative fuel value: %f", e.value)
}
