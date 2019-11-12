package main

import (
	"fmt"
	"github.com/arnobroekhof/rd2wgs84"
)

func convertToRD() {
	wgs84 := rd2wgs84.WGS84Coordinates{
		Lat: 52.37214383811702,
		Lon: 4.905597604352241,
	}

	rd := wgs84.ToRD()

	fmt.Println(rd.X) //122202
	fmt.Println(rd.Y) //487250
}
