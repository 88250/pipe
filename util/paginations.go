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

func Paginate(currentPageNum, pageSize, pageCount, windowsSize int) []int {
	ret := []int{}

	if pageCount < windowsSize {
		for i := 0; i < pageCount; i++ {
			ret = append(ret, i+1)
		}
	} else {
		first := currentPageNum + 1 - windowsSize/2
		if first < 1 {
			first = 1
		}
		if first+windowsSize > pageCount {
			first = pageCount - windowsSize + 1
		}
		for i := 0; i < windowsSize; i++ {
			ret = append(ret, first+i)
		}
	}

	return ret
}
