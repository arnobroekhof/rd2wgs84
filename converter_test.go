package rd2wgs84

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRDCoords_ToWGS84Coords(t *testing.T) {
	rdcoords := RDCoordinates{
		X: 122202,
		Y: 487250,
	}

	wgs84Coords := rdcoords.ToWGS84()
	assert.NotNil(t, wgs84Coords)
	assert.NotNil(t, wgs84Coords.Lat)
	assert.NotNil(t, wgs84Coords.Lon)

	assert.Equal(t, 52.37214383811702, wgs84Coords.Lat)
	assert.Equal(t, 4.905597604352241, wgs84Coords.Lon)
}

func TestWGS84Coords_toRDCoords(t *testing.T) {
	wgs84Coords := WGS84Coordinates{
		Lat: 52.37214383811702,
		Lon: 4.905597604352241,
	}

	rdCoords := wgs84Coords.toRD()
	assert.NotNil(t, rdCoords)
	assert.NotNil(t, rdCoords.X)
	assert.NotNil(t, rdCoords.Y)

	assert.Equal(t, 122202, rdCoords.X)
	assert.Equal(t, 487250, rdCoords.Y)

	t.Logf("Found X: %v", rdCoords.X)
	t.Logf("Found Y: %v", rdCoords.Y)
}
