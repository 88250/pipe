package util

import (
	"strings"
	"testing"
)

func TestRandImage(t *testing.T) {
	url := RandImage()
	if !strings.Contains(url, "img.hacpai.com") {
		t.Errorf(url)
	}
}

func TestRandImages(t *testing.T) {
	urls := RandImages(4)
	if 4 != len(urls) {
		t.Errorf("expected is [%d], actual is [%d]", 4, len(urls))
	}
}
