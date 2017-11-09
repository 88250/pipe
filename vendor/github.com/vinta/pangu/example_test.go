package pangu_test

import (
	"fmt"
	"github.com/vinta/pangu"
	"os"
)

func ExampleSpacingText() {
	s := pangu.SpacingText("所以,請問Jackey的鼻子有幾個?3.14個!")
	fmt.Println(s)
	// Output:
	// 所以, 請問 Jackey 的鼻子有幾個? 3.14 個!
}

func ExampleSpacingFile() {
	input := "_fixtures/test_file.txt"
	output := "_fixtures/test_file.pangu.txt"

	fw, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
	}
	defer fw.Close()

	pangu.SpacingFile(input, fw)
}
