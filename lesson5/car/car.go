package car

import (
	"mtsbank_golang/lesson5/navigator"
)

type Car struct {
	navigator       navigator.Navigator
	location        string
	address         string
	fuelTank        float64
	fuelConsumption float64
}

func (c *Car) Address() string {
	return c.address
}

func (c *Car) Location() string {
	return c.location
}

// двигаться по навигатору до конца маршрута
// если не хватило топлива, то вернуть false
func (c *Car) FollowNavigator() bool {

	for {
		if d, _ := c.navigator.NextDistance(); c.fuelTank-d*c.fuelConsumption >= 0 {
			if c.navigator.MoveNext() {
				c.fuelTank -= d * c.fuelConsumption
				c.location, _ = c.navigator.CurrentLocation()
				c.address, _ = c.navigator.CurrentAddress()
			} else {
				return true
			}
		} else {
			return false
		}
	}
}

func (c *Car) AddFuel(fuel float64) error {
	if fuel <= 0 {
		return &NegativeFuelValueError{fuel}
	}
	c.fuelTank += fuel
	return nil
}

func (c *Car) FuelTank() float64 {
	return c.fuelTank
}

func (c *Car) FuelConsumption() float64 {
	return c.fuelConsumption
}

func (c *Car) DistancePassed() float64 {
	return c.navigator.DistancePassed()
}

func NewCar(navigator navigator.Navigator, fuelTank float64, fuelConsumption float64) *Car {
	return &Car{navigator: navigator, fuelTank: fuelTank, fuelConsumption: fuelConsumption}
}
