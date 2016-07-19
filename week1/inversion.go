package main

import "fmt"

var r []int

func main() {
	//fmt.Println("start", a)
	a := mergesort([]int{7, 2, 6, 5, 4, 9, 8, 3})
	fmt.Println("done", a)
}

func mergesort(a []int) []int {
	if len(a) == 1 {
		return a
	} else if len(a) == 2 {
		if a[0] > a[1] {
			tmp := a[0]
			a[0] = a[1]
			a[1] = tmp
		}
		return a
	}

	return merge(mergesort(a[0:len(a)/2]), mergesort(a[len(a)/2:]))
}

func merge(b []int, c []int) []int {
	fmt.Println("B: ", b, "C: ", c)

	r = []int{}
	for i := 0; i < len(b); i++ {
		fmt.Println("B: ", b, "C: ", c)
		if b[i] > c[0] {
			r = append(r, c[0])
			c = append(c[1:])
		} else {
			r = append(r, b[i])
		}
	}
	fmt.Println(r)
	return r
}
