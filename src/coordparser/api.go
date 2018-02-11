package coordparser

import (
	"strconv"
	"strings"

	"github.com/pArkIngmate/src/boundingbox"
)

//GetBoxes returns the BoundingBoxes parsed from the file
func GetBoxes(lines []string) ([]boundingbox.BoundingBox, []error) {
	boxes := []boundingbox.BoundingBox{}
	errs := []error{}
	for _, line := range lines {
		coords, err := parseCoords(line)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		boxes = append(boxes, boundingbox.BoundingBox{
			Left:   coords[0],
			Top:    coords[1],
			Right:  coords[2],
			Bottom: coords[3],
		})
	}
	return boxes, errs
}

func parseCoords(line string) ([]int, error) {
	coords := strings.Split(line, "\t")
	nums := make([]int, len(coords))
	for idx, val := range coords {
		i, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		nums[idx] = i
	}
	//The order they will be returned is L, T, R, B
	return nums, nil
}
