package main

import "fmt"

func main2() {
	fmt.Println("result: ", mergesort([]int{7, 2, 6, 5, 4, 9, 8, 3}))
}

func mergesort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	return merge(len(a), mergesort(a[0:len(a)/2]), mergesort(a[len(a)/2:]))
}

func merge(n int, b []int, c []int) []int {
	var a []int
	i := 0
	j := 0

	for k := 0; k < n; {
		// check if out of bounds
		if j == len(c) {
			a = append(a, b[i:]...)
			break
		} else if i == len(b) {
			a = append(a, c[j:]...)
			break
		}

		// find which is less
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
