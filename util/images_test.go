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
)

func TestRandImage(t *testing.T) {
	url := RandImage()
	if !strings.Contains(url, CommunityFileURL) {
		t.Errorf(url)
	}
}

func TestRandImages(t *testing.T) {
	urls := RandImages(4)
	if 4 != len(urls) {
		t.Errorf("expected is [%d], actual is [%d]", 4, len(urls))
	}
}
