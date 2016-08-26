package mergesort

// Mergesort starts the recursive merging algorithm
func Mergesort(a []int) []int {
	if len(a) == 1 {
		return a
	}
	return merge(len(a), Mergesort(a[0:len(a)/2]), Mergesort(a[len(a)/2:]))
}

func merge(n int, b []int, c []int) []int {
	var a []int
	i, j := 0, 0

	for k := 0; k < n; {
		// check if out of bounds
		if i == len(b) {
			a = append(a, c[j:]...)
			break
		} else if j == len(c) {
			a = append(a, b[i:]...)
			break
		}

		// determine which is less
		if b[i] < c[j] {
			a = append(a, b[i])
			i++
		} else {
			a = append(a, c[j])
			j++
		}
	}
	return a
}
