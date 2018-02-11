package coordparser

import (
	"strconv"
	"strings"

	"github.com/pArkIngmate/src/boundingbox"
)

//GetBoxes returns the BoundingBoxes parsed from the file
func GetBoxes(lines chan string) (chan boundingbox.BoundingBox, chan error) {
	out := make(chan boundingbox.BoundingBox)
	errs := make(chan error)
	go func() {
		defer close(out)
		for line := range lines {
			go func(line string) {
				coords, err := parseCoords(line)
				if err != nil {
					errs <- err
					return
				}
				out <- boundingbox.BoundingBox{
					Left:   coords[0],
					Top:    coords[1],
					Right:  coords[2],
					Bottom: coords[3],
				}
			}(line)
		}
	}()
	return out, errs
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
