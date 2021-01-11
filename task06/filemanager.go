package task06

import (
	"math"
	"os"
	"path/filepath"
)

type FileInfo struct {
	FileName string
	Size     float64
}

func GetMaxSizeFileInDir(path string) (fileInfo FileInfo) {
	var maxSize int64

	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Size() > maxSize {
			maxSize = info.Size()
			fileInfo.FileName = info.Name()
			fileInfo.Size = math.Floor(float64(maxSize)/1073741824*100) / 100
		}

		return nil
	})

	return
}

func GetFileNamesInSizeRange(path string, minSizeInBytes int64, maxSizeInBytes int64) (filesInfo []FileInfo) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if info.Size() < minSizeInBytes || info.Size() > maxSizeInBytes {
			return nil
		}

		var fileInfo FileInfo
		fileInfo.FileName = info.Name()
		fileInfo.Size = math.Floor(float64(info.Size())/1073741824*100) / 100
		filesInfo = append(filesInfo, fileInfo)
		return nil
	})

	return
}
