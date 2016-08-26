package mergesort

import "testing"

func TestMergesort(t *testing.T) {
	merge := Mergesort([]int{7, 2, 6, 5, 4, 9, 8, 3})
	result := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for i, m := range merge {
		if m != result[i] {
			t.Error("error merging")
		}
	}
}

func TestMerge(t *testing.T) {
	merge := Merge([]int{1, 3, 5, 7}, []int{2, 4, 6, 8})
	result := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i, r := range result {
		if r != merge[i] {
			t.Error("error merging")
			break
		}
	}
}

// BenchmarkMergesort-4             2000000               886 ns/op
func BenchmarkMergesort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mergesort([]int{7, 2, 6, 5, 4, 9, 8, 3})
	}
}

// BenchmarkMerge-4                10000000               178 ns/op
func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Merge([]int{1, 3, 5, 7}, []int{2, 4, 6, 8})
	}
}
