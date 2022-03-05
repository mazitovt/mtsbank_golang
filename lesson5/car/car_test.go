package car

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mtsbank_golang/lesson5/distance/path"
	"mtsbank_golang/lesson5/distance/points"
	"mtsbank_golang/lesson5/geocoding"
	"mtsbank_golang/lesson5/mock"
	"mtsbank_golang/lesson5/navigator"
	"testing"
)

func TestNavigator(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGeocoder := mock.NewMockGeocoder(ctrl)

	p1, _ := points.NewPointOnSphereFromDegrees(55.7539, 37.6208)
	p2, _ := points.NewPointOnSphereFromDegrees(55.7539, 30.3146)
	p3, _ := points.NewPointOnSphereFromDegrees(59.9398, 30.3146)
	p4, _ := points.NewPointOnSphereFromDegrees(59.9398, 37.6208)

	mockGeocoder.EXPECT().ReverseGeocode(*p1).Return(geocoding.GeocodeData{Point: *p1, Address: "first point address", PostalCode: "1"}, nil)
	mockGeocoder.EXPECT().ReverseGeocode(*p2).Return(geocoding.GeocodeData{Point: *p2, Address: "second point address", PostalCode: "2"}, nil)
	mockGeocoder.EXPECT().ReverseGeocode(*p3).Return(geocoding.GeocodeData{Point: *p3, Address: "third point address", PostalCode: "3"}, nil)
	mockGeocoder.EXPECT().ReverseGeocode(*p4).Return(geocoding.GeocodeData{Point: *p4, Address: "fourth point address", PostalCode: "4"}, nil)

	pathOnSphere := path.NewPathOnSphere(6731, []points.PointOnSphere{*p1, *p2, *p3, *p4})

	navi, _ := navigator.NewPlanetNavigator(mockGeocoder, *pathOnSphere)
	//navi, _ := navigator.NewPathNavigator(pathOnSphere)

	car1 := NewCar(navi, 400, 1)

	f := car1.FollowNavigator()
	fmt.Printf("loc: %v ,addr: %v, fuel: %v,  follow: %v\n", car1.Location(), car1.Address(), car1.FuelTank(), f)

	_ = car1.AddFuel(500)
	f = car1.FollowNavigator()
	fmt.Printf("loc: %v ,addr: %v, fuel: %v,  follow: %v\n", car1.Location(), car1.Address(), car1.FuelTank(), f)

	_ = car1.AddFuel(500)
	f = car1.FollowNavigator()
	fmt.Printf("loc: %v ,addr: %v, fuel: %v,  follow: %v\n", car1.Location(), car1.Address(), car1.FuelTank(), f)

	_ = car1.AddFuel(100)
	f = car1.FollowNavigator()
	fmt.Printf("loc: %v ,addr: %v, fuel: %v,  follow: %v\n", car1.Location(), car1.Address(), car1.FuelTank(), f)

	assert.Equal(t, car1.FuelTank(), 95.73264660895683)
}
