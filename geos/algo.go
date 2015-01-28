package geos

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lgeos_c
#include <geos_c.h>

*/
import "C"

const (
	ORIENTATION_CCW Orientation = -1
	ORIENTATION_CW  Orientation = 1
	ORIENTATION_CL  Orientation = 0
)

type Orientation int

func CalcOrientation(ax float64, ay float64, bx float64, by float64, px float64, py float64) Orientation {
	val := C.GEOSOrientationIndex_r(ctxHandle, C.double(ax), C.double(ay), C.double(bx), C.double(by), C.double(px), C.double(py))
	return Orientation(val)
}
