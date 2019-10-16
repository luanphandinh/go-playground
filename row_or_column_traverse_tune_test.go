package go_tuning_examples

import "testing"

const DIMENSION = 10000
var arr [DIMENSION][DIMENSION]int

func rowTraverse() int {
	count := 0
	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {
			count += arr[i][j]
		}
	}

	return count
}

func colTraverse() int {
	count := 0
	for i := 0; i < DIMENSION; i++ {
		for j := 0; j < DIMENSION; j++ {
			count += arr[j][i]
		}
	}

	return count
}

func BenchmarkRowTraverse(b *testing.B) {
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
