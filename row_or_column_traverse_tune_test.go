package go_tuning_examples

import (
	"fmt"
	"testing"
)

const ROWS = 10000
const COLS = 10000

var arr [][]int64

func initTraverseTestArray() {
	fmt.Println("Init array")
	arr = make([][]int64, ROWS)
	for i := 0; i < ROWS; i++ {
		arr[i] = make([]int64, COLS)
	}

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			arr[i][j] = int64(i + j)
		}
	}
	fmt.Println("Init array done.")
}

func rowTraverse() int64 {
	count := int64(0)
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			count += arr[row][col]
		}
	}

	return count
}

func colTraverse() int64 {
	count := int64(0)
	for col := 0; col < COLS; col++ {
		for row := 0; row < ROWS; row++ {
			count += arr[row][col]
		}
	}

	return count
}

func BenchmarkRowTraverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rowTraverse()
	}
}

func BenchmarkColTraverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		colTraverse()
	}
}
