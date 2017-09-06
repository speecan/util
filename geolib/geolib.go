package geolib

import "math"

// Deg2Rad converts a degree to randian
func Deg2Rad(degree float64) float64 {
	return degree * math.Pi / 180
}

// Rad2Deg converts a randian to degree
func Rad2Deg(radian float64) float64 {
	return radian * 180 / math.Pi
}

// TileNumToLonLat returns lon lat
// form zxy tile number
func TileNumToLonLat(z, x, y int) (float64, float64) {
	xtile := float64(x) + 0.5
	ytile := float64(y) + 0.5
	n := math.Pow(2, float64(z))
	lon := (xtile/n)*360 - 180
	lat := Rad2Deg(math.Atan(math.Sinh(math.Pi * (1 - (2*ytile)/n))))
	return lon, lat
}

// LonLatToTileNum returns z, x, y int
// form zoom and lon lat
func LonLatToTileNum(zoom int, lon, lat float64) (z, x, y int) {
	xtile := math.Floor(((lon + 180) / 360) * math.Pow(2, float64(zoom)))
	ytile := math.Floor((1 - math.Log(math.Tan(Deg2Rad(lat))+1/math.Cos(Deg2Rad(lat)))/math.Pi) / 2 * math.Pow(2, float64(zoom)))
	z = zoom
	x = int(xtile)
	y = int(ytile)
	return
}
