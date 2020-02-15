// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (c) 2017-present, b3log.org
//
// Pipe is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package util

import (
	"testing"
)

func TestPaginate(t *testing.T) {
	pageNums := paginate(1, 15, 99, 20)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i, val := range pageNums {
		if val != expected[i] {
			t.Errorf("exptected is [%d] at index [%d], actual is [%d]", expected[i], i, val)
		}
	}

	pageNums = paginate(50, 15, 99, 20)
	expected = []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
	for i, val := range pageNums {
		if val != expected[i] {
			t.Errorf("exptected is [%d] at index [%d], actual is [%d]", expected[i], i, val)
		}
	}
}
