// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
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
	"math/rand"
	"strconv"
	"time"
	"strings"
)

// ImageSize returns image URL of Qiniu image processing style with the specified width and height.
func ImageSize(imageURL string, width, height int) string {
	if strings.Contains(imageURL, "imageView") || !strings.Contains(imageURL, "img.hacpai.com") {
		return imageURL
	}

	return imageURL + "?imageView2/1/w/" + strconv.Itoa(width) + "/h/" + strconv.Itoa(height) + "/interlace/1/q/100"
}

// RandImage returns an image URL randomly for article thumbnail.
// https://github.com/88250/bing
func RandImage() string {
	min := time.Date(2017, 11, 04, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min

	return time.Unix(sec, 0).Format("https://img.hacpai.com/bing/20060102.jpg")
}

// RandImages returns random image URLs.
func RandImages(n int) []string {
	var ret []string

	i := 0
	for {
		if i >= n*5 {
			break
		}

		url := RandImage()
		if !contains(url, ret) {
			ret = append(ret, url)
		}

		if len(ret) >= n {
			return ret
		}

		i++
	}

	return ret
}

func contains(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}

	return false
}
