package coordparser_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pArkIngmate/src/coordparser"
)

func Test_GetBoxes_ReturnsBoundingBoxes(t *testing.T) {
	testStrings := []string{}

	for i := 0; i < 100; i++ {
		testStrings = append(testStrings, (string(i) + "\t"))
	}

	boxes, errs := coordparser.GetBoxes(testStrings)
	i := 0
	for _, box := range boxes {
		fmt.Printf(
			"\n\nL: %d\tT: %d\tR: %d\tB: %d",
			box.Left,
			box.Top,
			box.Right,
			box.Bottom,
		)
		assert.Equal(t, i, box.Left)
		i++
		assert.Equal(t, i, box.Top)
		i++
		assert.Equal(t, i, box.Right)
		i++
		assert.Equal(t, i, box.Bottom)
		i++
	}
	for _, err := range errs {
		fmt.Println(err.Error())
	}
}
