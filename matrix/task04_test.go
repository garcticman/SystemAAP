package matrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	matrix := Matrix{
		{5, 2, 3, 5, 6},
		{4, 6, 2, 7, 1},
	}

	matrixRequaire := Matrix{
		{5, 4},
		{2, 6},
		{3, 2},
		{5, 7},
		{6, 1},
	}

	result, err := matrix.Transpose()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result, matrixRequaire) {
		t.Fatal()
	}
}

func TestMultiply(t *testing.T) {
	matrixA := Matrix{
		{2, -3},
		{4, -6},
	}
	matrixB := Matrix{
		{9, -6},
		{6, -4},
	}
	matrixRequaire := Matrix{
		{0, 0},
		{0, 0},
	}

	matrixResult, err := Multiply(matrixA, matrixB)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(matrixResult, matrixRequaire) {
		t.Fatal()
	}

	matrixA = Matrix{
		{-2, 1},
		{5, 4},
	}
	matrixB = Matrix{
		{3},
		{-1},
	}
	matrixRequaire = Matrix{
		{-7},
		{11},
	}

	matrixResult, err = Multiply(matrixA, matrixB)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(matrixResult, matrixRequaire) {
		t.Fatal()
	}
}