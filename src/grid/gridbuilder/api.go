package gridbuilder

import (
	"github.com/pArkIngmate/src/boundingbox"
	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
	"github.com/pArkIngmate/src/grid/gridbuilder/gridbuildertypes"
)

//GridBuilder creates the grid to estimate lines of parking rows
type GridBuilder struct {
	Boxes []*boundingbox.BoundingBox
}

//New returns a new gridbuilder
func New(boxes []*boundingbox.BoundingBox) *GridBuilder {
	return &GridBuilder{
		Boxes: boxes,
	}
}

func getCenters(boxes []boundingbox.BoundingBox) []boundingboxtypes.Coordinate {
	out := []boundingboxtypes.Coordinate{}
	for _, val := range boxes {
		out = append(out, val.Center())
	}
	return out
}

//GetLines returns the lines created by the boxes
func (g *GridBuilder) GetLines() []*gridbuildertypes.Line {
	lines := []*gridbuildertypes.Line{}
	for i, val := range g.Boxes {
		for j := i; j < len(g.Boxes); j++ {
			lines = append(lines, gridbuildertypes.NewLine(val, g.Boxes[j]))
		}
	}
	return lines
}

func findIntersection(l *gridbuildertypes.Line, b *boundingbox.BoundingBox) bool {
	for x := b.Left; x < b.Right && x < l.End.X; x++ {
		step := x - l.Start.X
		val := l.Slope*float64(step) + float64(l.Start.Y)
		if val > float64(b.Bottom) && val < float64(b.Top) {
			return true
		}
	}
	return false
}

//DetermineIntersects finds all the intersecting bounding boxes in a grid for the given line
func (g *GridBuilder) DetermineIntersects(l *gridbuildertypes.Line) []*boundingbox.BoundingBox {
	intersections := []*boundingbox.BoundingBox{}
	for _, box := range g.Boxes {
		if findIntersection(l, box) {
			intersections = append(intersections, box)
		}
	}
	return intersections
}

//GetConfidence gets a rough idea of how accurate the line is
func (g *GridBuilder) GetConfidence(l *gridbuildertypes.Line) float64 {
	return float64(len(l.Intersects)) / float64(len(g.Boxes))
}
