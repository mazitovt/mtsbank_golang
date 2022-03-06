package points

import "fmt"

type InvalidLatitudeError struct {
	lat float64
}

func (e *InvalidLatitudeError) Error() string {
	return fmt.Sprintf("широта точки '%f' должна лежать в пределах от -90 до 90", e.lat)
}

type InvalidLongitudeError struct {
	lon float64
}

func (e *InvalidLongitudeError) Error() string {
	return fmt.Sprintf("долгота точки '%f' должна лежать в пределах от -180 до 180", e.lon)
}
