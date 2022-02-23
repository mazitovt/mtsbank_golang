package path

import "fmt"

type InvalidNumberOfPointForDistance struct {
	numArgs int
}

func (e *InvalidNumberOfPointForDistance) Error() string {
	return fmt.Sprintf("расчет длины пути:\n\tнеобходимо минимумЖ 2 точки\n\tполучено: '%f'", e.numArgs)
}
