// SPDX-License-Identifier: Apache-2.0
// Copyright Bernardo Soares

package matrix

/*

Matrix multiplication using the naive algorithm.

// Matrices are measured by Rows × Columns.
// If you are multiplying Matrix A by Matrix B:
// Matrix A has dimensions: m x n
// Matrix B has dimensions: n x p
// Because the inner numbers match (n and n), you can safely multiply them.
// The resulting matrix will always take the outer dimensions: m × p

*/

// Types

// Plain simple representation of a matrix using
// an array of arrays.
type Matrix [][]int

func (m *Matrix) Multiply(other Matrix) (Matrix, error) {
	return nil, nil
}

// Flatten transforms a Matrix into a FlatMatrix
func (m *Matrix) Flatten() FlatMatrix {
	var totalElements int

	for i := range *m {
		totalElements += len((*m)[i])
	}

	flat := FlatMatrix{
		rows:    len(*m),
		columns: totalElements / len(*m),
		matrix:  make([]int, totalElements),
	}

	var position int
	for i := range len(*m) {
		copy(flat.matrix[position:], (*m)[i])
		position += len((*m)[i])
	}

	return flat
}

// Flattened version of a matrix, where rows and columns
// sit on the same array
type FlatMatrix struct {
	rows    int
	columns int
	matrix  []int
}

func (m *FlatMatrix) Multiply(other FlatMatrix) (Matrix, error) {
	return nil, nil
}

func (m *FlatMatrix) UnFlatten() Matrix {
	result := make(Matrix, m.rows)

	for i := range m.rows {
		start := i * m.columns
		end := start + m.columns

		result[i] = m.matrix[start:end]
	}

	return result
}
