package main

import (
	"fmt"
	"github.com/arnobroekhof/rd2wgs84"
)

func main() {

	rd := rd2wgs84.RDCoordinates{
		X: 122202,
		Y: 487250,
	}

	wgs84 := rd.ToWGS84()

	fmt.Println(wgs84.Lat)
	fmt.Println(wgs84.Lon)

}
