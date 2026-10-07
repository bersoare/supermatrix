package matrix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// dont optimize my code away!!!!!
var matrixResult Matrix
var flatMatrixResult FlatMatrix

func TestMatrixMultipl(t *testing.T) {
	t.SkipNow()
	one := Matrix{
		{1, 2, 3},
		{5, 6, 7},
		{9, 8, 7},
		{5, 3, 2},
	}

	two := Matrix{
		{12, -5},
		{24, 13},
		{29, -999},
	}

	expected := Matrix{
		{147, -2976},
		{407, -6940},
		{503, -6934},
		{190, -1984},
	}

	t.Run("matrixMultiplicationOneWorker", func(t *testing.T) {
		result, err := one.Multiply(two)
		require.NoError(t, err)
		require.Equal(t, expected, result)
	})
}

func BenchmarkMatrixUtils(b *testing.B) {
	matrixA := generateMatrix(10, 30)
	matrixB := generateFlatMatrix(30, 10)

	b.Logf("%s: %d x %d (%d) items", b.Name(), len(matrixA), matrixB.rows, len(matrixA)*matrixB.columns)

	b.Run("flatten", func(b *testing.B) {
		b.ResetTimer()

		var result FlatMatrix

		for range b.N {
			result = matrixA.Flatten()
		}

		b.StopTimer()
		require.NotEmpty(b, result)
		require.Equal(b, len(matrixA), result.rows)
		require.Equal(b, len(matrixA[0]), result.columns)
		flatMatrixResult = result
	})

	b.Run("unflatten", func(b *testing.B) {
		b.ResetTimer()

		var result Matrix

		for range b.N {
			result = matrixB.UnFlatten()
		}

		b.StopTimer()
		require.NotEmpty(b, result)
		require.Equal(b, matrixB.rows, len(result))
		require.Equal(b, matrixB.columns, len(result[0]))
		matrixResult = result
	})
}
