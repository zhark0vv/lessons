package benchmark

import (
	"testing"
)

func BenchmarkLinearSearch(b *testing.B) {
	size := 10000000
	target := size - 1
	arr := generateSortedSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(arr, target)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	size := 10000000
	target := size - 1
	arr := generateSortedSlice(size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(arr, target)
	}
}
