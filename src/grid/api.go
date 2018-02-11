package grid

import (
	"fmt"
	"log"

	"github.com/pArkIngmate/src/boundingbox"
	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
	"github.com/pArkIngmate/src/grid/gridbuilder"
	"github.com/pArkIngmate/src/grid/gridbuilder/gridbuildertypes"
)

//BuildGrid returns the gridbuilder object and the lines
func BuildGrid(boxes []*boundingbox.BoundingBox) (*gridbuilder.GridBuilder, []*gridbuildertypes.Line) {
	gb := gridbuilder.New(boxes)
	lines := gb.GetLines()
	for _, line := range lines {
		line.Intersects = gb.DetermineIntersects(line)
		line.Confidence = gb.GetConfidence(line)
		log.Println((line.Confidence))
	}
	return gb, lines
}

//GetGapPoints returns the potential parking lot spots
func GetGapPoints(l *gridbuildertypes.Line) []*boundingboxtypes.Coordinate {
	coords := []*boundingboxtypes.Coordinate{}
	for i := range l.Intersects {
		fmt.Println("intersect found")
		if i < len(l.Intersects)-1 {
			if l.Intersects[i].Right < l.Intersects[i+1].Left {
				xMid := (l.Intersects[i].Right + l.Intersects[i+1].Left) / 2
				step := xMid - l.Start.X
				yMid := float64(l.Start.Y) + float64(l.Slope)*float64(step)
				coord := &boundingboxtypes.Coordinate{X: int(xMid), Y: int(yMid)}
				log.Printf("Empty spot found: %d, %d\n", coord.X, coord.Y)
				coords = append(coords, coord)
			}
		}
	}
	return coords
}
