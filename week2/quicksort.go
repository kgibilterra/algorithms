package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var comparisons int

func main() {

	//A := []int{3, 9, 8, 4, 6, 10, 2, 5, 7, 1}
	f, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("omg can't read ur file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	data := []int{}
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		data = append(data, x)
	}
	fmt.Println("Starting Array:", data[:10])

	A := partitionArray(data, 0, len(data))
	fmt.Println("Final Array:", A[:10])
	fmt.Println("comparisons:", comparisons)

}

func partitionArray(A []int, l int, r int) []int {
	//comparisons += len(A) - 1
	if len(A) < 2 {
		return A
	}

	// don't use anything when pivot is the first element
	//p := A[l]

	// use this when pivot is the last element
	A[r-1], A[l] = A[l], A[r-1]
	p := A[l]

	// "median" pivot
	// median := len(A) / 2
	// if len(A)%2 == 0 {
	// 	median--
	// }
	// pivots := []int{A[l], A[r-1], A[median]}
	// p := pivots[0]
	// if pivots[1] < pivots[2] && pivots[1] < pivots[0] {
	// 	p = pivots[1]
	// 	A[r-1], A[l] = A[l], A[r-1]
	// } else if pivots[2] < pivots[1] && pivots[2] < pivots[0] {
	// 	p = pivots[2]
	// 	A[median], A[l] = A[l], A[median]
	// }

	// partition Array
	i := l + 1
	for j := l + 1; j < r; j++ {
		comparisons++
		if A[j] < p {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[l], A[i-1] = A[i-1], A[l]

	// call partition on left and right sides of array
	partitionArray(A[:i-1], 0, len(A[:i-1]))
	partitionArray(A[i:], 0, len(A[i:]))

	return A
}

func selectPivot(A []int) int {
	return 0
}
