package pangu_test

import (
	"github.com/vinta/pangu"
	"testing"
)

func BenchmarkSpacingText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pangu.SpacingText("所以,請問Jackey的鼻子有幾個?3.14個!")
	}
}

func BenchmarkSpacingFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleSpacingFile()
	}
}
