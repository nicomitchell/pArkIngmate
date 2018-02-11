package boundingboxiface

import "github.com/pArkIngmate/src/boundingbox/boundingboxtypes"

//Box is the interface that determines the behavior of the BoundingBox struct
type Box interface {
	Center() boundingboxtypes.Coordinate
	Area() int
	Height() int
	Width() int
}
