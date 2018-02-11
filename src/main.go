package main

import (
	"bufio"
	"fmt"
	"image/png"
	"io"
	"log"
	"os"

	"github.com/pArkIngmate/src/imgwriter/imageeditor"

	"github.com/pArkIngmate/src/boundingbox"
	"github.com/pArkIngmate/src/imgwriter"

	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
	"github.com/pArkIngmate/src/coordparser"
)

func main() {
	f, err := os.Open("/home/nicolasmitchell/go/src/github.com/pArkIngmate/bounding_box.txt")
	if err != nil {
		panic(fmt.Sprintf("File open error: %s", err.Error()))
	}
	defer f.Close()
	lines := getLines(f)
	boxes, errs := coordparser.GetBoxes(lines)
	for _, err := range errs {
		log.Printf("Parser err: " + err.Error())
	}
	centers := getCenters(boxes)

	for _, val := range centers {
		fmt.Printf("%d, %d\n", val.X, val.Y)
	}
	imgFile, err := os.Open("/home/nicolasmitchell/go/src/github.com/darknet/predictions.png")
	if err != nil {
		panic(fmt.Sprintf("File open error: %s", err.Error()))
	}
	defer imgFile.Close()
	img, err := png.Decode(imgFile)
	if err != nil {
		panic(fmt.Sprintf("PNG decoding error: %s", err.Error()))
	}
	outFile, err := os.Create("/home/nicolasmitchell/go/src/github.com/pArkIngmate/image_out.jpg")
	if err != nil {
		panic(fmt.Sprintf("File open error: %s", err))
	}
	defer outFile.Close()

	//pts := getCrossPixels(centers)
	imgEditor := imageeditor.New(img)
	iWriter := imgwriter.New(centers, nil, imgEditor, outFile)
	//iWriter.SetAllPoints()
	iWriter.SetAllPointsCircles(10)
	err = iWriter.Save()
	if err != nil {
		panic(fmt.Sprintf("Save Error: %s", err.Error()))
	}
	log.Println("Successfully saved file")
}

func getLines(f io.Reader) []string {
	lines := []string{}
	r := bufio.NewScanner(f)
	var line []byte
	for {
		check := r.Scan()
		if !check {
			break
		}
		line = r.Bytes()
		fmt.Println(string(line))
		lines = append(lines, string(line))
	}
	return lines
}

func getCenters(boxes []boundingbox.BoundingBox) []boundingboxtypes.Coordinate {
	centers := []boundingboxtypes.Coordinate{}
	for _, box := range boxes {
		centers = append(centers, box.Center())
	}
	return centers
}

func getCrossPixels(pts []boundingboxtypes.Coordinate) []boundingboxtypes.Coordinate {
	var out []boundingboxtypes.Coordinate
	for _, val := range pts {
		for i := -15; i < 16; i++ {
			out = append(out, boundingboxtypes.Coordinate{X: (val.X + i), Y: val.Y})
			out = append(out, boundingboxtypes.Coordinate{X: val.X, Y: (val.Y + i)})

		}
	}
	return out
}
