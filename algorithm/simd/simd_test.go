package simd

import (
	"fmt"
	"math/rand/v2"
	"testing"
	"time"
)

func generateRandomMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Float64() // 產生 0.0 到 1.0 之間的隨機浮點數
		}
	}
	return matrix
}

func isEqual(matrix1, matrix2 [][]float64) bool {
	rows := len(matrix1)
	cols := len(matrix1[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMatrixMultiplicationPerformance(t *testing.T) {
	sizes := []int{2048}

	for _, size := range sizes {
		// 为每个尺寸创建一个子测试单元
		testName := fmt.Sprintf("MatrixSize_%dx%d", size, size)
		t.Run(testName, func(t *testing.T) {
			// 1. 准备数据
			matrix1 := generateRandomMatrix(size)
			matrix2 := generateRandomMatrix(size)
			var result1, result2, result3, result4 [][]float64

			// 2. 测试标准方法并计时
			start1 := time.Now()
			result1 = MatrixMultiply(matrix1, matrix2)
			duration1 := time.Since(start1)
			t.Logf("标准方法 (Standard)   运行时间: %s", duration1)

			// 3. 测试转置优化方法并计时
			start2 := time.Now()
			result2 = MatrixMultiplyTrans(matrix1, matrix2)
			duration2 := time.Since(start2)
			t.Logf("转置优化 (Transpose) 运行时间: %s", duration2)

			// 4. 并行计算方法
			start3 := time.Now()
			result3 = ParelleMultiply(matrix1, matrix2)
			duration3 := time.Since(start3)
			t.Logf("并行计算 (Tiled) 运行时间: %s", duration3)

			// 5.分块计算
			start4 := time.Now()
			result4 = MatrixMultiplyBlock(matrix1, matrix2)
			duration4 := time.Since(start4)
			t.Logf("分块并行计算 (Tiled) 运行时间: %s", duration4)

			// 5. 结果验证
			if !isEqual(result1, result2) || !isEqual(result1, result3) || !isEqual(result2, result2) || !isEqual(result4, result1) {
				t.Errorf("计算结果不一致！")
			}
		})
	}
}
