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
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Pagination represents pagination info.
type Pagination struct {
	CurrentPageNum  int    `json:"currentPageNum"`
	PageSize        int    `json:"pageSize"`
	PageCount       int    `json:"pageCount"`
	WindowSize      int    `json:"windowSize"`
	RecordCount     int    `json:"recordCount"`
	PageNums        []int  `json:"pageNums"`
	NextPageNum     int    `json:"nextPageNum"`
	PreviousPageNum int    `json:"previousPageNum"`
	FirstPageNum    int    `json:"firstPageNum"`
	LastPageNum     int    `json:"lastPageNum"`
	PageURL         string `json:"pageURL"`
}

// GetPage returns paging parameter.
func GetPage(c *gin.Context) int {
	ret, _ := strconv.Atoi(c.Query("p"))
	if 1 > ret {
		ret = 1
	}

	return ret
}

// NewPagination creates a new pagination with the specified current page num, page size, window size and record count.
func NewPagination(currentPageNum, pageSize, windowSize, recordCount int) *Pagination {
	pageCount := int(math.Ceil(float64(recordCount) / float64(pageSize)))

	previousPageNum := currentPageNum - 1
	if 1 > previousPageNum {
		previousPageNum = 0
	}
	nextPageNum := currentPageNum + 1
	if nextPageNum > pageCount {
		nextPageNum = 0
	}

	pageNums := paginate(currentPageNum, pageSize, pageCount, windowSize)
	firstPageNum := 0
	lastPageNum := 0
	if 0 < len(pageNums) {
		firstPageNum = pageNums[0]
		lastPageNum = pageNums[len(pageNums)-1]
	}

	return &Pagination{
		CurrentPageNum:  currentPageNum,
		NextPageNum:     nextPageNum,
		PreviousPageNum: previousPageNum,
		PageSize:        pageSize,
		PageCount:       pageCount,
		WindowSize:      windowSize,
		RecordCount:     recordCount,
		PageNums:        pageNums,
		FirstPageNum:    firstPageNum,
		LastPageNum:     lastPageNum,
	}
}

func paginate(currentPageNum, pageSize, pageCount, windowSize int) []int {
	var ret []int

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
