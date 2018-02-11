package boundingbox

import (
	"github.com/pArkIngmate/src/boundingbox/boundingboxiface"
	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
)

//BoundingBox contains the coordinates of one bounded object
type BoundingBox struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

//New returns a new bounding box that implements the interface
func New(l, t, r, b int) boundingboxiface.Box {
	return &BoundingBox{
		Left:   l,
		Right:  r,
		Top:    t,
		Bottom: b,
	}
}

//Center returns the point at the center of the Bounding Box
func (b *BoundingBox) Center() (coord boundingboxtypes.Coordinate) {
	return boundingboxtypes.Coordinate{
		X: int((b.Left + b.Right) / 2),
		Y: int((b.Top + b.Bottom) / 2),
	}
}

//Area returns the area of the bounding box
func (b *BoundingBox) Area() int {
	return b.Height() * b.Width()
}

//Height returns the height of the box
func (b *BoundingBox) Height() int {
	return b.Top - b.Bottom
}

//Width returns the width of the box
func (b *BoundingBox) Width() int {
	return b.Right - b.Left
}
