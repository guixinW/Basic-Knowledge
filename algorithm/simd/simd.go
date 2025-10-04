package simd

import (
	"runtime"
	"sync"
)

func MatrixMultiply(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	ret := make([][]float64, len(matrix1))
	for i := range ret {
		ret[i] = make([]float64, len(matrix2[0]))
	}
	rowsA := len(matrix1)
	colsA := len(matrix1[0])
	colsB := len(matrix2[0])
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			var sum float64
			for k := 0; k < colsA; k++ {
				sum += matrix1[i][k] * matrix2[k][j]
			}
			ret[i][j] = sum
		}
	}
	return ret
}

func transpose(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return matrix
	}
	rows := len(matrix)
	cols := len(matrix[0])
	transposed := make([][]float64, cols)
	for i := range transposed {
		transposed[i] = make([]float64, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

func MatrixMultiplyTrans(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	ret := make([][]float64, len(matrix1))
	for i := range ret {
		ret[i] = make([]float64, len(matrix2[0]))
	}
	rowsA := len(matrix1)
	colsA := len(matrix1[0])
	colsB := len(matrix2[0])
	matrix2T := transpose(matrix2)
	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			var sum float64
			for k := 0; k < colsA; k++ {
				sum += matrix1[i][k] * matrix2T[j][k]
			}
			ret[i][j] = sum
		}
	}
	return ret
}

func ParelleMultiply(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	rows1 := len(matrix1)
	cols1 := len(matrix1[0])
	cols2 := len(matrix2[0])

	result := make([][]float64, rows1)
	for i := range result {
		result[i] = make([]float64, cols2)
	}

	// 转置第二个矩阵
	matrix2T := transpose(matrix2)

	numWorkers := runtime.NumCPU()
	rowsPerWorker := rows1 / numWorkers
	if rowsPerWorker == 0 {
		rowsPerWorker = 1
	}

	var wg sync.WaitGroup

	for start := 0; start < rows1; start += rowsPerWorker {
		end := start + rowsPerWorker
		if end > rows1 {
			end = rows1
		}

		wg.Add(1)
		go func(startRow, endRow int) {
			defer wg.Done()
			for i := startRow; i < endRow; i++ {
				for j := 0; j < cols2; j++ {
					var sum float64
					for k := 0; k < cols1; k++ {
						sum += matrix1[i][k] * matrix2T[j][k] // 使用转置矩阵
					}
					result[i][j] = sum
				}
			}
		}(start, end)
	}
	wg.Wait()
	return result
}

func MatrixMultiplyBlock(matrix1 [][]float64, matrix2 [][]float64) [][]float64 {
	ret := make([][]float64, len(matrix1))
	for i := range ret {
		ret[i] = make([]float64, len(matrix2[0]))
	}

	rowsA := len(matrix1)
	colsA := len(matrix1[0])
	colsB := len(matrix2[0])

	for i := 0; i < rowsA; i++ {
		for j := 0; j < colsB; j++ {
			var sum float64
			k := 0

			// 循环展开，每次处理4个元素
			for ; k < colsA-3; k += 4 {
				sum += matrix1[i][k]*matrix2[k][j] +
					matrix1[i][k+1]*matrix2[k+1][j] +
					matrix1[i][k+2]*matrix2[k+2][j] +
					matrix1[i][k+3]*matrix2[k+3][j]
			}

			for ; k < colsA; k++ {
				sum += matrix1[i][k] * matrix2[k][j]
			}

			ret[i][j] = sum
		}
	}

	return ret
}
