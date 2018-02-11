package imageeditor_test

import (
	"image"
	"testing"

	"github.com/pArkIngmate/src/imageeditor"
	"github.com/stretchr/testify/assert"
)

func Test_New_ImplementsInterface(t *testing.T) {
	assert.Implements(t, (*image.Image)(nil), imageeditor.New(nil))
}
