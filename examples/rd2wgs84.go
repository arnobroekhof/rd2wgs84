package main

import (
	"fmt"
	rd2wgs84 "github.com/arnobroekhof/rd2wsg84"
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
