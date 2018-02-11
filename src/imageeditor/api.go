package imageeditor

import (
	"image"
	"image/color"

	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
)

//ChangedImage wraps the image.Image struct and allows us to change pixel color when we save a new image
type ChangedImage struct {
	changed map[boundingboxtypes.Coordinate]color.Color
	Img     image.Image
}

//New returns a new ChangedImage which contains a regular image
func New(img image.Image) image.Image {
	return &ChangedImage{
		Img: img,
	}
}

//At will tell the image encoder the color of a pixel - the "changed" values contain all the pixels whose colors are not
//the same as in the original image
func (i *ChangedImage) At(x, y int) color.Color {
	if val, check := i.changed[boundingboxtypes.Coordinate{X: x, Y: y}]; check {
		return val
	}
	return i.Img.At(x, y)
}

//Bounds returns the image bounds
func (i *ChangedImage) Bounds() image.Rectangle {
	return i.Img.Bounds()
}

//ColorModel returns the image's color model
func (i *ChangedImage) ColorModel() color.Model {
	return i.Img.ColorModel()
}
