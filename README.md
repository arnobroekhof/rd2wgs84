# Rijksdriehoek 2 WGS84 Converter

The RD (Rijks-Driehoek) system is the coordinate system used by the Dutch geographical service. 
Conversion to and from WGS84 (World Geodetic System; latitude and longitude pairs) is not entirely straightforward. 
The conversion software is based on the formulas and coefficients in this (Dutch) document (contains a summary in English). 
This library works for coordinates within the Netherlands (and some area outside it). 
Highest precision is achieved West of 6.0Â°E. You can enter either RD-coordinates (in meters) or WGS84-coordinates in degrees using a decimal point. 


# example

```go

package main

import (
  "github.com/arnobroekhof/rd2wgs84"
  "fmt"
)

rdcoords := RDCoords{
		X: 122202,
		Y: 487250,
	}

	wgs84Coords := rdcoords.ToWGS84Coords()


```