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
			ip:       "192.168.0.1",
			expected: true,
		},
		{
			ip:       "0.0.0.0",
			expected: true,
		},
	}

	for _, grid := range grids {
		actual := IsIPv4(grid.ip)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestIsCellphoneOfChina(t *testing.T) {
	grids := []struct {
		cellphone string
		expected  bool
	}{
		{
			cellphone: "12323434334",
			expected:  false,
		},
		{
			cellphone: "13717995000",
			expected:  true,
		},
		{
			cellphone: "+8613717995000",
			expected:  true,
		},
		{
			cellphone: "+86 13717995000",
			expected:  true,
		},
		{
			cellphone: "+86-13717995000",
			expected:  true,
		},
		{
			cellphone: "+86 -13717995000",
			expected:  false,
		},
		{
			cellphone: "+86- 13717995000",
			expected:  false,
		},
		{
			cellphone: "+86 - 13717995000",
			expected:  false,
		},
	}

	for _, grid := range grids {
		actual := IsCellphoneOfChina(grid.cellphone)
		assert.Equal(t, grid.expected, actual)
	}
}

func TestIsTelephoneOfChina(t *testing.T) {
	grids := []struct {
		telephone string
		expected  bool
	}{
		{
			telephone: "010-10086",
			expected:  true,
		},
		{
			telephone: "10086",
			expected:  true,
		},
		{
			telephone: "96069",
			expected:  true,
		},
		{
			telephone: "0350-346321d",
			expected:  false,
		},
		{
			telephone: "0350-346321",
			expected:  true,
		},
		{
			telephone: "0350-3463241",
			expected:  true,
		},
		{
			telephone: "010-92832839",
			expected:  true,
		},
		{
			telephone: "010-02832839",
			expected:  false,
		},
	}

	for _, grid := range grids {
		actual := IsTelephoneOfChina(grid.telephone)
		assert.Equal(t, grid.expected, actual)
	}
}
