package task05

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScalarProduct(t *testing.T) {
	vectors := []Vector{
		{2, -5},
		{-1, 0},
	}

	result, err := ScalarProduct(vectors)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	if !reflect.DeepEqual(result, float64(-2)) {
		t.Fatal()
	}

	vectors = []Vector{
		{2, 4, -1, 6, -4, 0},
		{5, 11, -5, 6, 2, 1},
	}

	result, err = ScalarProduct(vectors)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	if !reflect.DeepEqual(result, float64(87)) {
		t.Fatal()
	}
}