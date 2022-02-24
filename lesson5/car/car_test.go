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
	mockGeocoder.EXPECT().ReverseGeocode(*p2).Return(geocoding.GeocodeData{Point: *p2, Address: "second point address error", PostalCode: "2"}, nil)
	mockGeocoder.EXPECT().ReverseGeocode(*p3).Return(geocoding.GeocodeData{Point: *p3, Address: "third point address", PostalCode: "3"}, nil)
	mockGeocoder.EXPECT().ReverseGeocode(*p4).Return(geocoding.GeocodeData{Point: *p4, Address: "fourth point address", PostalCode: "4"}, nil)

	pathOnSphere := path.NewPathOnSphere(6731, []points.PointOnSphere{*p1, *p2, *p3, *p4})

	navi, _ := navigator.NewPlanetNavigator(mockGeocoder, *pathOnSphere)

	car1 := NewCar(navi, 400, 1)
	f := car1.FollowNavigator()
	fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)

	_ = car1.AddFuel(500)

	f = car1.FollowNavigator()

	fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)
	_ = car1.AddFuel(500)

	f = car1.FollowNavigator()
	fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)
	_ = car1.AddFuel(100)

	f = car1.FollowNavigator()
	fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)

	assert.Equal(t, car1.FuelTank(), 95.73264660895683)

	//info, err := navi.PathInfo(*point1, *point2)
	//info1, err1 := navi.PathInfo(*point1, *point2)
	//fmt.Println(info1, err1)
	//assert.Nil(t, err)
	//fmt.Println(info.PlaceStart().Point)
	//assert.Equal(t, info.PlaceStart().Point, *point1)
	//assert.Equal(t, info.PlaceFinish().Point, *point2)
	//assert.Equal(t, info.PlaceFinish().PostalCode, "111321")
	//assert.Equal(t, info.PlaceStart().PostalCode, "123321")

}

func TestPath(t *testing.T) {

}
