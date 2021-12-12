package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegularMatchString(t *testing.T) {
	grids := []struct {
		patten   string
		str      string
		expected bool
	}{
		{
			patten:   `no`,
			str:      `/usr/bin/which: no hello in (/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin)`,
			expected: true,
		},
		{
			patten:   `not found`,
			str:      `hello not found`,
			expected: true,
		},
		{
			patten:   `file or directory not found`,
			str:      `hello not found`,
			expected: false,
		},
	}

	for _, grid := range grids {
		actual, err := RegularMatchString(grid.patten, grid.str)
		assert.NoError(t, err)
		assert.Equal(t, grid.expected, actual)
	}
}
