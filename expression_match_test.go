// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
