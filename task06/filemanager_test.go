package task06

import (
	"fmt"
	"testing"
)

func TestGetMaxSizeFileInDir(t *testing.T) {
	fmt.Println(GetMaxSizeFileInDir(`D:\канал`))
}

func TestGetFileNamesInSizeRange(t *testing.T) {
	fmt.Println(GetFileNamesInSizeRange(`D:\канал`, 2061694668, 2061694668))
}
