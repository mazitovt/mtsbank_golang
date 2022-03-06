package main

import (
	"fmt"
	"mtsbank_golang/tests_mocks/distance/path"
	"mtsbank_golang/tests_mocks/distance/points"
	"mtsbank_golang/tests_mocks/navigator"
)

func main() {
	p1, _ := points.NewPointOnSphereFromDegrees(55.7539, 37.6208)
	p2, _ := points.NewPointOnSphereFromDegrees(55.7539, 30.3146)
	p3, _ := points.NewPointOnSphereFromDegrees(59.9398, 30.3146)
	p4, _ := points.NewPointOnSphereFromDegrees(59.9398, 37.6208)

	pathOnSphere := path.NewPathOnSphere(6731, []points.PointOnSphere{*p1, *p2, *p3, *p4})

	navi, _ := navigator.NewPathNavigator(pathOnSphere)

	for navi.MoveNext() {
		v, e := navi.CurrentLocation()
		fmt.Println(e)
		fmt.Println(v)
		fmt.Println("left: ", navi.DistanceLeft())
		fmt.Println("passed: ", navi.DistancePassed())
	}

	//p1, _ := points.NewPointOnSphereFromDegrees(55.7539, 37.6208)
	//p2, _ := points.NewPointOnSphereFromDegrees(55.7539, 30.3146)
	//p3, _ := points.NewPointOnSphereFromDegrees(59.9398, 30.3146)
	//p4, _ := points.NewPointOnSphereFromDegrees(59.9398, 37.6208)
	//
	//path1 := path.NewPathOnSphere(6731, []points.PointOnSphere{*p1, *p2, *p3, *p4})
	//
	//navi, _ := navigator.NewPlanetNavigator(geocoding.NewGeocoder(), *path1)
	////navi, _ := navigator.NewPathNavigator(path1)
	//
	//car1 := car.NewCar(navi, 500, 1)
	//f := car1.FollowNavigator()
	//fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)
	//
	//_ = car1.AddFuel(500)
	//
	//f = car1.FollowNavigator()
	//
	//fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)
	//_ = car1.AddFuel(500)
	//
	//f = car1.FollowNavigator()
	//fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)
	//_ = car1.AddFuel(100)
	//
	//f = car1.FollowNavigator()
	//fmt.Printf("car.location: %v , fuel: %v,  follow: %v\n", car1.Location(), car1.FuelTank(), f)

}

//p1, _ := points.NewPointOnSphereFromDegrees(55.7539, 37.6208)
//p2, _ := points.NewPointOnSphereFromDegrees(55.7539, 30.3146)
//p3, _ := points.NewPointOnSphereFromDegrees(59.9398, 30.3146)
//p4, _ := points.NewPointOnSphereFromDegrees(59.9398, 37.6208)
//
//pathOnSphere := path.NewPathOnSphere(6731, []points.PointOnSphere{*p1, *p2, *p3, *p4})
//
//navi, _ := navigator.NewPathNavigator(pathOnSphere)
//
//for navi.MoveNext() {
//v, _ := navi.CurrentLocation()
//fmt.Println(v)
//fmt.Println("left: ", navi.DistanceLeft())
//fmt.Println("passed: ", navi.DistancePassed())
//}
//
//p2d1 := points.NewPoint2d(1, 1)
//p2d2 := points.NewPoint2d(2, 1)
//p2d3 := points.NewPoint2d(2, 2)
//p2d4 := points.NewPoint2d(1, 2)
//
//path2d := path.NewPath2d([]points.Point2d{*p2d1, *p2d2, *p2d3, *p2d4})
//
//navi2d, _ := navigator.NewPathNavigator(path2d)
//
//for navi2d.MoveNext() {
//v, _ := navi2d.CurrentLocation()
//fmt.Println(v)
//fmt.Println("left: ", navi2d.DistanceLeft())
//fmt.Println("passed: ", navi2d.DistancePassed())
//}
//
//p3d1 := points.NewPoint3d(1, 1, 1)
//p3d2 := points.NewPoint3d(2, 1, 2)
//p3d3 := points.NewPoint3d(2, 2, 3)
//p3d4 := points.NewPoint3d(1, 2, 4)
//
//path3d := path.NewPath3d([]points.Point3d{*p3d1, *p3d2, *p3d3, *p3d4})
//
//navi3d, _ := navigator.NewPathNavigator(path3d)
//
//for navi3d.MoveNext() {
//v, _ := navi3d.CurrentLocation()
//fmt.Println(v)
//fmt.Println("left: ", navi3d.DistanceLeft())
//fmt.Println("passed: ", navi3d.DistancePassed())
//}
