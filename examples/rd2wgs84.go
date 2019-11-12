package main

import (
	"fmt"
	"github.com/arnobroekhof/rd2wgs84"
)

func convertToWGS84() {

	rd := rd2wgs84.RDCoordinates{
		X: 122202,
		Y: 487250,
	}

	wgs84 := rd.ToWGS84()

	fmt.Println(wgs84.Lat) //52.37214383811702
	fmt.Println(wgs84.Lon)  //4.905597604352241

}
