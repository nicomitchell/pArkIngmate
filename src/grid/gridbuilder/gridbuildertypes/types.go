package gridbuildertypes

import (
	"github.com/pArkIngmate/src/boundingbox"
	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
)

//Line represents a graphed line of pixels - should intersect with center points of bounding boxes
type Line struct {
	Start, End        boundingboxtypes.Coordinate
	Confidence, Slope float64
	StartBox, EndBox  *boundingbox.BoundingBox
	Intersects        []*boundingbox.BoundingBox
}

//NewLine creates a new Line object from the centerpoints and boxes (they should be associated)
func NewLine(sBox, eBox *boundingbox.BoundingBox) *Line {
	start := sBox.Center()
	end := eBox.Center()
	slope := (float64(end.Y - start.Y)) / float64((end.X / start.X))
	return &Line{
		Start:    start,
		End:      end,
		Slope:    slope,
		StartBox: sBox,
		EndBox:   eBox,
	}
}
