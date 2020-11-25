package task05

import "errors"

type Vector []float64

func ScalarProduct(vectors []Vector) (opResult float64, err error) {
	if len(vectors) <= 1 {
		return 0, nil
	}

	for i := 1; i < len(vectors); i++ {
		if len(vectors[i]) != len(vectors[i-1]) {
			return 0, errors.New("vectors length are not equal")
		}
	}

	tmpChan := make(chan float64, len(vectors[0]))
	for i := 0; i < len(vectors[0]); i++ {
		i := i
		go func() {
			tmp := float64(1)
			for j := 0; j < len(vectors); j++{
				tmp *= vectors[j][i]
			}
			tmpChan <- tmp
		}()
	}
	for i := 0; i < len(vectors[0]); i++ {
		opResult += <-tmpChan
	}

	return
}