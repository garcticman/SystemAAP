package tasks

import "strings"

var fileExtensions = []string{
	".jpg",
	".png",
	".jpeg",
	".bmp",
}

func getImageNames(allNames []string) (imageNames []string) {
	for _, name := range allNames {
		for _, extension := range fileExtensions {
			if strings.Contains(name, extension) {
				imageNames = append(imageNames, name)
			}
		}
	}
	return
}

func getMapFromNameSlice(allNames []string) (result map[string]int) {
	for _, name := range allNames {
		result[name] = len(name)
	}
	return
}