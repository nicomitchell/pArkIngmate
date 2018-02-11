package imgwriter

import (
	"image/color"
	"io"
	"math"

	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
	"github.com/pArkIngmate/src/imgwriter/imageeditor"
)

//ImageWriter is the struct used to modify and save images
type ImageWriter struct {
	img         *imageeditor.ChangedImage
	greenPoints []boundingboxtypes.Coordinate
	redPoints   []boundingboxtypes.Coordinate
	w           io.Writer
}

//New returns a new ImageWriter object - redpoints come first
func New(rpts []boundingboxtypes.Coordinate, gpts []boundingboxtypes.Coordinate, img *imageeditor.ChangedImage, w io.Writer) *ImageWriter {
	return &ImageWriter{
		img:         img,
		redPoints:   rpts,
		greenPoints: gpts,
		w:           w,
	}
}

//SetAllPoints sets all of the green and red points to be green and red in the image
func (i *ImageWriter) SetAllPoints() {
	for _, val := range i.greenPoints {
		i.img.Set(val, color.RGBA{10, 150, 50, 200})
	}
	for _, val := range i.redPoints {
		i.img.Set(val, color.RGBA{200, 20, 20, 150})
	}
}

//SetAllPointsCircles draws circles around each of the points
func (i *ImageWriter) SetAllPointsCircles(r int) {
	for _, val := range i.greenPoints {
		for xOffset := 0 - r; xOffset < r+1; xOffset++ {
			for yOffset := 0 - r; yOffset < r+1; yOffset++ {
				if math.Sqrt(math.Pow(float64(xOffset), 2.0)+math.Pow(float64(yOffset), 2)) < float64(r) {
					pt := boundingboxtypes.Coordinate{X: val.X + xOffset, Y: val.Y + yOffset}
					i.img.Set(pt, color.RGBA{10, 150, 50, 200})

				}
			}
		}
	}
	for _, val := range i.redPoints {
		for xOffset := 0 - r; xOffset < r+1; xOffset++ {
			for yOffset := 0 - r; yOffset < r+1; yOffset++ {
				if math.Sqrt(math.Pow(float64(xOffset), 2.0)+math.Pow(float64(yOffset), 2)) < float64(r) {
					pt := boundingboxtypes.Coordinate{X: val.X + xOffset, Y: val.Y + yOffset}
					i.img.Set(pt, color.RGBA{200, 20, 20, 150})
				}
			}
		}
	}
}

//Save saves the file using the writer
func (i *ImageWriter) Save() error {
	return i.img.Save(i.w)
}
