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

type Pagination struct {
	CurrentPageNum int   `json:"currentPageNum"`
	PageSize       int   `json:"pageSize"`
	PageCount      int   `json:"pageCount"`
	WindowSize     int   `json:"windowSize"`
	RecordCount    int   `json:"recordCount"`
	PageNums       []int `json:"pageNums"`
}

func NewPagination(currentPageNum, pageSize, pageCount, windowSize, recordCount int) *Pagination {
	return &Pagination{
		CurrentPageNum: currentPageNum,
		PageSize:       pageSize,
		PageCount:      pageCount,
		WindowSize:     windowSize,
		RecordCount:    recordCount,
		PageNums:       paginate(currentPageNum, pageSize, pageCount, windowSize),
	}
}

func paginate(currentPageNum, pageSize, pageCount, windowSize int) []int {
	ret := []int{}

	if pageCount < windowSize {
		for i := 0; i < pageCount; i++ {
			ret = append(ret, i+1)
		}
	} else {
		first := currentPageNum + 1 - windowSize/2
		if first < 1 {
			first = 1
		}
		if first+windowSize > pageCount {
			first = pageCount - windowSize + 1
		}
		for i := 0; i < windowSize; i++ {
			ret = append(ret, first+i)
		}
	}

	return ret
}
