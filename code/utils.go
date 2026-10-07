package matrix

import "math/rand"

// generateMatrix creates a matrix with the requested amount of columns
// and rows.
func generateMatrix(columns, rows int) Matrix {
	result := make(Matrix, rows)
	for i := range result {
		result[i] = make([]int, columns)

		for ii := range result[i] {
			result[i][ii] = rand.Intn(7676)
		}
	}

	return result
}

// generateFlatMatrix creates a matrix with the requested amount of columns
// and rows.
func generateFlatMatrix(columns, rows int) FlatMatrix {
	result := FlatMatrix{
		columns: columns,
		rows:    rows,
		matrix:  make([]int, columns*rows),
	}

	for i := range result.matrix {
		result.matrix[i] = rand.Intn(7676)
	}

	return result
}
