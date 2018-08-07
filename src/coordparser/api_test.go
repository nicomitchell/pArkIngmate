package coordparser_test

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pArkIngmate/src/coordparser"
)

func Test_GetBoxes_ReturnsBoundingBoxes(t *testing.T) {
	var s string
	for i := 0; i < 100; i += 4 {
		s = fmt.Sprintf("%s\n%s", s, fmt.Sprintf("%d\t%d\t%d\t%d", i, i+1, i+2, i+3))
	}
	lines := strings.Split(s, "\n")
	boxes, errs := coordparser.GetBoxes(lines)
	i := 0
	for _, box := range boxes {
		log.Printf(
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
