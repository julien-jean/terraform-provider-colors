package colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomColor(t *testing.T) {
	expectedColors := []string{
		"blue",
		"red",
		"green",
		"white",
		"black",
	}
	color := RandomColor()

	assert.Contains(t, expectedColors, color.Name)
}
