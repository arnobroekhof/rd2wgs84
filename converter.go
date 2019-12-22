package rd2wgs84

import "math"

const (
	X0   = 155000
	Y0   = 463000
	lat0 = 52.15517
	lon0 = 5.387206
)

func NewWSG84Coordinates(lat, lon float64) WGS84Coordinates {
	return WGS84Coordinates{
		Lat: lat,
		Lon: lon,
	}
}

type WGS84Coordinates struct {
	Lat float64
	Lon float64
}

func NewRDCoordinates(x, y float64) RDCoordinates {
	return RDCoordinates{
		X: x,
		Y: y,
	}
}

type RDCoordinates struct {
	X float64
	Y float64
}

func (r RDCoordinates) ToWGS84() WGS84Coordinates {
	dX := (r.X - X0) * math.Pow(10, -5)
	dY := (r.Y - Y0) * math.Pow(10, -5)

	sumN := (3235.65389 * dY) +
		(-32.58297 * math.Pow(dX, 2)) +
		(-0.2475 * math.Pow(dY, 2)) +
		(-0.84978 * math.Pow(dX, 2) * dY) +
		(-0.0655 * math.Pow(dY, 3)) +
		(-0.01709 * math.Pow(dX, 2) * math.Pow(dY, 2)) +
		(-0.00738 * dX) +
		(0.0053 * math.Pow(dX, 4)) +
		(-0.00039 * math.Pow(dX, 2) * math.Pow(dY, 3)) +
		(0.00033 * math.Pow(dX, 4) * dY) +
		(-0.00012 * dX * dY)

	sumE := (5260.52916 * dX) +
		(105.94684 * dX * dY) +
		(2.45656 * dX * math.Pow(dY, 2)) +
		(-0.81885 * math.Pow(dX, 3)) +
		(0.05594 * dX * math.Pow(dY, 3)) +
		(-0.05607 * math.Pow(dX, 3) * dY) +
		(0.01199 * dY) +
		(-0.00256 * math.Pow(dX, 3) * math.Pow(dY, 2)) +
		(0.00128 * dX * math.Pow(dY, 4)) +
		(0.00022 * math.Pow(dY, 2)) +
		(-0.00022 * math.Pow(dX, 2)) +
		(0.00026 * math.Pow(dX, 5))

	lat := lat0 + (sumN / 3600)
	lon := lon0 + (sumE / 3600)

	return WGS84Coordinates{Lat: lat,
		Lon: lon}
}

func (w WGS84Coordinates) ToRD() RDCoordinates {
	var x float64
	var y float64

	rp := []float64{0, 1, 2, 0, 1, 3, 1, 0, 2}
	rq := []float64{1, 1, 1, 3, 0, 1, 3, 2, 3}
	rpq := []float64{190094.945, -11832.228, -114.221, -32.391, -0.705, -2.340, -0.608, -0.008, 0.148}

	sp := []float64{1, 0, 2, 1, 3, 0, 2, 1, 0, 1}
	sq := []float64{0, 2, 0, 2, 0, 1, 2, 1, 4, 4}
	spq := []float64{309056.544, 3638.893, 73.077, -157.984, 59.788, 0.433, -6.439, -0.032, 0.092, -0.054}

	lat := 0.36 * (w.Lat - lat0)
	lon := 0.36 * (w.Lon - lon0)

	for r := range rpq {
		x = x + (rpq[r] * math.Pow(lat, rp[r]) * math.Pow(lon, rq[r]))
	}

	for s := range spq {
		y = y + (spq[s] * math.Pow(lat, sp[s]) * math.Pow(lon, sq[s]))
	}

	return RDCoordinates{
		X: math.Round(X0 + x),
		Y: math.Round(Y0 + y),
	}
}
