package grid

import (
	"github.com/pArkIngmate/src/boundingbox"
	"github.com/pArkIngmate/src/grid/gridbuilder"
	"github.com/pArkIngmate/src/grid/gridbuilder/gridbuildertypes"
)

//BuildGrid returns the gridbuilder object and the lines
func BuildGrid(boxes []*boundingbox.BoundingBox) (*gridbuilder.GridBuilder, []*gridbuildertypes.Line) {
	gb := gridbuilder.New(boxes)
	lines := gb.GetLines()
	for _, line := range lines {
		line.Intersects := gb.DetermineIntersects(line)
		line.Confidence := gb.GetConfidence(line)
	}
}
