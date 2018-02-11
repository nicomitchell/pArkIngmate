package boundingbox

import "github.com/pArkIngmate/src/boundingbox/boundingboxiface"

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
func (b *BoundingBox) Center() (x, y int) {
	return int((b.Left + b.Right) / 2), int((b.Top + b.Bottom) / 2)
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
