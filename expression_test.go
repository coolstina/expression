package expression

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIPv4(t *testing.T) {
	grids := []struct {
		ip       string
		expected bool
	}{
		{
			ip:       "192.168.0.1:8000",
			expected: true,
		},
		{
			ip:       "192.168.0.1",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsIPv4(grid.ip)
		assert.Equal(t, grid.expected, actual)
	}
}
