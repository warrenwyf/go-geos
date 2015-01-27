package geos

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	s := createCoordSeq(2, 3)

	s.setX(1, 1)
	s.setY(1, 2)
	s.setZ(1, 3)

	size := s.getSize()
	if size != 2 {
		t.Errorf("Error: getSize() returns wrong result %d, want 2", size)
	}

	var val float64

	val = s.getX(1)
	if val != 1 {
		t.Errorf("Error: getX() returns wrong result %q, want 1", val)
	}

	val = s.getY(1)
	if val != 2 {
		t.Errorf("Error: getY() returns wrong result %q, want 2", val)
	}

	val = s.getZ(1)
	if val != 3 {
		t.Errorf("Error: getZ() returns wrong result %q, want 3", val)
	}
}

func TestToCoords(t *testing.T) {
	s := createCoordSeq(2, 3)

	coords := s.toCoords()
	if len(coords) != 2 {
		t.Errorf("Error: toCoords() returns wrong result")
	}

	coordZs := s.toCoordZs()
	if len(coordZs) != 2 {
		t.Errorf("Error: toCoordZs() returns wrong result")
	}

}
