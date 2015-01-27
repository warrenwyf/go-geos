package geos

import (
	"fmt"
)

type Coord struct {
	X, Y float64
}

type CoordZ struct {
	X, Y, Z float64
}

func (p *Coord) Equals(p2 *Coord) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *CoordZ) Equals(p2 *CoordZ) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}

func (p *Coord) ToString() string {
	return fmt.Sprintf("[%f, %f]", p.X, p.Y)
}

func (p *CoordZ) ToString() string {
	return fmt.Sprintf("[%f, %f, %f]", p.X, p.Y, p.Z)
}
