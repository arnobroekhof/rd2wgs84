package rd2wgs84

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRDCoords_ToWGS84Coords(t *testing.T) {
	rdcoords := NewRDCoordinates(122202, 487250)

	wgs84Coords := rdcoords.ToWGS84()
	assert.NotNil(t, wgs84Coords)
	assert.NotNil(t, wgs84Coords.Lat)
	assert.NotNil(t, wgs84Coords.Lon)

	assert.Equal(t, 52.37214383811702, wgs84Coords.Lat)
	assert.Equal(t, 4.905597604352241, wgs84Coords.Lon)
}

func TestWGS84Coords_toRDCoords(t *testing.T) {
	wgs84Coords := NewWSG84Coordinates(52.37214383811702, 4.905597604352241)

	rdCoords := wgs84Coords.ToRD()
	assert.NotNil(t, rdCoords)
	assert.NotNil(t, rdCoords.X)
	assert.NotNil(t, rdCoords.Y)

	assert.Equal(t, float64(122202), rdCoords.X)
	assert.Equal(t, float64(487250), rdCoords.Y)

	t.Logf("Found X: %v", rdCoords.X)
	t.Logf("Found Y: %v", rdCoords.Y)
}

func BenchmarkRDCoordinates_ToWGS84(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rdcoords := NewRDCoordinates(122202, 487250)
		rdcoords.ToWGS84()
	}
}

func BenchmarkWGS84Coordinates_ToRD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wgs84Coords := NewWSG84Coordinates(52.37214383811702, 4.905597604352241)
		wgs84Coords.ToRD()
	}
}
