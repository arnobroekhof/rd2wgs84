# Rijksdriehoek 2 WGS84 Converter

The RD (Rijks-Driehoek) system is the coordinate system used by the Dutch geographical service. 
Conversion to and from WGS84 (World Geodetic System; latitude and longitude pairs) is not entirely straightforward. 
The conversion software is based on the formulas and coefficients in this (Dutch) document (contains a summary in English). 
This library works for coordinates within the Netherlands (and some area outside it). 
Highest precision is achieved West of 6.0Â°E. You can enter either RD-coordinates (in meters) or WGS84-coordinates in degrees using a decimal point. 


## examples

convert WGS84 to Rijksdriehoek

```go

package main

import (
	"fmt"
	"github.com/arnobroekhof/rd2wgs84"
)

func main() {
    wgs84 := rd2wgs84.NewWSG84Coordinates(52.37214383811702, 4.905597604352241)
    rd := wgs84.ToRD()
    
    fmt.Println(rd.X) //122202
    fmt.Println(rd.Y) //487250

}

```

convert Rijksdriehoek coordinates to WGS84

```go
package main

import (
	"fmt"
	"github.com/arnobroekhof/rd2wgs84"
)

func main() {
    rd := rd2wgs84.NewRDCoordinates(122202, 487250)
    wgs84 := rd.ToWGS84()
    
    fmt.Println(wgs84.Lat) //52.37214383811702
    fmt.Println(wgs84.Lon) //4.905597604352241

}
```
