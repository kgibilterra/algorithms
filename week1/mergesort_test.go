package mergesort

import "testing"

func TestMergesort(t *testing.T) {
	merge := Mergesort([]int{7, 2, 6, 5, 4, 9, 8, 3})
	result := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for i, m := range merge {
		if m != result[i] {
			t.Error("error merging!")
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
