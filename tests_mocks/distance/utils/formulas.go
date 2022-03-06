package utils

import "math"

func RadiansToDegrees(radian float64) float64 {
	return radian * 180 / math.Pi
}
func DegreesToRadians(degree float64) float64 {
	return degree * math.Pi / 180
}

func ShortestDistanceTwoPoints2D(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

func ShortestDistanceTwoPoints3D(x1, y1, z1, x2, y2, z2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2) + math.Pow(z1-z2, 2))
}

func ShortestDistanceTwoPointsOnSphere(r, lat1, lon1, lat2, lon2 float64) float64 {
	return 2 * r * math.Asin(math.Sqrt(math.Pow(math.Sin((lat2-lat1)/2), 2)+math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lon2-lon1)/2), 2)))
}
