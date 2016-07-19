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

	return merge3(len(a), mergesort(a[0:len(a)/2]), mergesort(a[len(a)/2:]))
}

func merge(b []int, c []int) []int {
	fmt.Println("B: ", b, "C: ", c)

	r = []int{}
	for i := 0; i < len(b); i++ {
		fmt.Println("B: ", b[i], "C: ", c[0])
		if b[i] > c[0] {
			r = append(r, c[0])
			c = append(c[:0], c[1:]...)
			fmt.Println("DELETE C? ", c)
		} else {
			r = append(r, b[i])
		}
		fmt.Println("R: ", r)
	}
	fmt.Println(r)
	return r
}

func merge2(b []int, c []int) []int {
	fmt.Println("B: ", b, "C: ", c)

	r = []int{}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(c); j++ {
			fmt.Println("B: ", b[i], "C: ", c[0])
			if b[i] > c[j] {
				r = append(r, c[j])
			} else {
				r = append(r, b[i])
				fmt.Println("START C: ", c)
				c = append(c[:j], c[j:]...)
				fmt.Println("END C: ", c)
				break
			}

		}
	}
	fmt.Println(r)
	return r
}

func merge3(n int, b []int, c []int) []int {
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
