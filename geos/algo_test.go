package geos

import (
	"testing"
)

func TestCalcOrientation(t *testing.T) {

	if CalcOrientation(0, 0, 10, 0, 5, 1) != ORIENTATION_CW {
		t.Errorf("Error: CalcOrientation(0, 0, 10, 0, 5, 1) error")
	}

	if CalcOrientation(0, 0, 10, 0, 5, -1) != ORIENTATION_CCW {
		t.Errorf("Error: CalcOrientation(0, 0, 10, 0, 5, -1) error")
	}

	if CalcOrientation(0, 0, 10, 0, 5, 0) != ORIENTATION_CL {
		t.Errorf("Error: CalcOrientation(0, 0, 10, 0, 5, 0) error")
	}

}
