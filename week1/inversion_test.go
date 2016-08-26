package mergesort

import "testing"

func TestCombine(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 6, 8}
	c := Combine(a, b)
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, r := range result {
		if r != c[i] {
			t.Error("error combining!")
			break
		}
	}
}

func TestInversionDivide(t *testing.T) {
	a := []int{1, 3, 6, 8, 2, 9, 4, 7, 5}
	b := InversionDivide(a)
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, r := range result {
		if r != b[i] {
			t.Error("error in inversion divide!")
			break
		}
	}
}

// BenchmarkInversionDivide-4       2000000              1053 ns/op
func BenchmarkInversionDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InversionDivide([]int{1, 3, 6, 8, 2, 9, 4, 7, 5})
	}
}

// BenchmarkCombine-4               5000000               274 ns/op
func BenchmarkCombine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Combine([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8})
	}
}
