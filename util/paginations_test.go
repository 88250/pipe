// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"testing"
)

func TestPaginate(t *testing.T) {
	pageNumbs := Paginate(1, 15, 99, 20)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, val := range pageNumbs {
		if val != expected[i] {
			t.Errorf("exptected is [%d] at index [%d], actual is [%d]", expected[i], i, val)
		}
	}

	pageNumbs = Paginate(50, 15, 99, 20)
	expected = []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
	for i, val := range pageNumbs {
		if val != expected[i] {
			t.Errorf("exptected is [%d] at index [%d], actual is [%d]", expected[i], i, val)
		}
	}
}
