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
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CommunityFileURL is the community file service URL.
const CommunityFileURL = "https://b3logfile.com"

// ImageSize returns image URL of Qiniu image processing style with the specified width and height.
func ImageSize(imageURL string, width, height int) string {
	if !Uploaded(imageURL) || strings.Contains(imageURL, "imageView") || strings.Contains(imageURL, ".gif") {
		return imageURL
	}

	return imageURL + "?imageView2/1/w/" + strconv.Itoa(width) + "/h/" + strconv.Itoa(height) + "/interlace/1/q/100"
}

// Uploaded checks whether the specified URL has uploaded.
func Uploaded(url string) bool {
	return strings.HasPrefix(url, CommunityFileURL) || strings.HasPrefix(url, "https://img.hacpai.com")
}

// RandImage returns an image URL randomly for article thumbnail.
// https://github.com/88250/bing
func RandImage() string {
	min := time.Date(2017, 11, 04, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min

	return CommunityFileURL + "/bing/" + time.Unix(sec, 0).Format("20060102") + ".jpg"
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
