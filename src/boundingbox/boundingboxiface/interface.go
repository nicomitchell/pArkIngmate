package boundingboxiface

//Box is the interface that determines the behavior of the BoundingBox struct
type Box interface {
	Center() (int, int)
	Area() int
	Height() int
	Width() int
}
