// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
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

func TestIsEmail(t *testing.T) {
	grids := []struct {
		email    string
		expected bool
	}{
		{
			email:    "wu.shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "a@tom.com",
			expected: true,
		},
		{
			email:    "a.b@tom.com",
			expected: true,
		},
		{
			email:    "a.helloshaohua@tom.com.cn",
			expected: true,
		},
		{
			email:    "a@foxmail.com",
			expected: true,
		},
		{
			email:    "wu.shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "wu-shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "wu_shaohua@foxmail.com",
			expected: true,
		},
		{
			email:    "hello.@foxmail.com",
			expected: false,
		},
		{
			email:    "hello.world....@foxmail.com",
			expected: false,
		},
		{
			email:    "helle..@foxmail.com",
			expected: false,
		},
	}

	for _, grid := range grids {
		actual := IsEmail(grid.email)
		assert.Equal(t, grid.expected, actual,
			"want: %t, but got: %t, email: %+v\n", grid.expected, actual, grid.email)
	}
}
