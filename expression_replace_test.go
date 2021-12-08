package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegularReplaceMatch(t *testing.T) {
	grids := []struct {
		patten   string
		original string
		replace  string
		expected string
	}{
		{
			patten:   `\{\{\.NAME\}\}`,
			original: "The {{.NAME}} Application starting...",
			replace:  `CallMe`,
			expected: "The CallMe Application starting...",
		},
	}

	for _, grid := range grids {
		actual, err := RegularReplaceMatch(grid.patten, grid.original, grid.replace)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, actual)
	}
}
