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
	"strings"
	"testing"
	"time"
)

func TestRandImage(t *testing.T) {
	originalLocal := time.Local
	defer func() { time.Local = originalLocal }()
	for _, offset := range []int{0, -12 * 60 * 60, 14 * 60 * 60} {
		time.Local = time.FixedZone("test", offset)
		for i := 0; i < 10000; i++ {
			url := RandImage()
			prefix := CommunityFileURL + "/bing/"
			if !strings.HasPrefix(url, prefix) || !strings.HasSuffix(url, ".jpg") {
				t.Fatalf("unexpected image URL: %s", url)
			}
			date := strings.TrimSuffix(strings.TrimPrefix(url, prefix), ".jpg")
			if _, err := time.Parse("20060102", date); err != nil {
				t.Fatalf("invalid image date in %s: %v", url, err)
			}
			if date < "20171104" || date > "20231220" {
				t.Fatalf("image date outside archive (offset %d): %s", offset, url)
			}
		}
	}
}

func TestRandImages(t *testing.T) {
	urls := RandImages(4)
	if 4 != len(urls) {
		t.Errorf("expected is [%d], actual is [%d]", 4, len(urls))
	}
}
