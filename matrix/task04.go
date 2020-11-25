package matrix

import (
	"errors"
)

type Matrix [][]float64

func (m Matrix) IsEmpty() bool {
	return m == nil || len(m) == 0
}

func (m Matrix) LinesCount() int {
	return len(m)
}

func (m Matrix) ColumnsCount() int {
	if m.IsEmpty() {
		return 0
	}

	return len(m[0])
}

func (m *Matrix) InitColumns(columnsCount int) {
	for i := range *m {
		(*m)[i] = make([]float64, columnsCount)
	}
}

func (m Matrix) Transpose() (transposedMatrix Matrix, err error) {
	if m.IsEmpty() {
		return nil, errors.New("nil matrix")
	}

	transposedMatrix = make(Matrix, m.ColumnsCount())
	transposedMatrix.InitColumns(m.LinesCount())

	for i := 0; i < m.LinesCount(); i++ {
		for j := 0; j < m.ColumnsCount(); j++ {
			transposedMatrix[j][i] = m[i][j]
		}
	}

	return transposedMatrix, nil
}

func Multiply(matrixA, matrixB Matrix) (opResult Matrix, err error) {
	if matrixA.IsEmpty() || matrixB.IsEmpty() {
		return nil, errors.New("matrix is empty")
	}
	if matrixA.ColumnsCount() != matrixB.LinesCount() {
		return nil, errors.New("matrixA columns not equal matrixB lines")
	}

	opResult = make(Matrix, matrixB.LinesCount())
	opResult.InitColumns(matrixB.ColumnsCount())

	for i := 0; i < matrixB.ColumnsCount(); i++ {
		for k := 0; k < matrixA.LinesCount(); k++ {
			for j := 0; j < matrixA.ColumnsCount(); j++ {
				opResult[k][i] += matrixA[k][j] * matrixB[j][i]
			}
		}
	}

	return opResult, nil
}