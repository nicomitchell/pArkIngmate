package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pArkIngmate/src/boundingbox"

	"github.com/pArkIngmate/src/boundingbox/boundingboxtypes"
	"github.com/pArkIngmate/src/coordparser"
)

func main() {
	f, err := os.Open("/home/nicolasmitchell/go/src/github.com/pArkIngmate/bounding_box.txt")
	if err != nil {
		log.Printf("File open error: %s", err.Error())
	}
	lines := getLines(f)
	boxes, errs := coordparser.GetBoxes(lines)
	for _, err := range errs {
		log.Printf("Parser err: " + err.Error())
	}
	centers := getCenters(boxes)

	for _, val := range centers {
		fmt.Printf("%d, %d\n", val.X, val.Y)
	}
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
}
