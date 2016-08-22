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
	fmt.Println("length of data:", len(data), "first 10 elements:", data[:10])
	fmt.Println("Starting Array:", data)

	A := partitionArray(data)
	fmt.Println("Final Array:", A)
	fmt.Println("comparisons:", comparisons)

}

func partitionArray(A []int) []int {
	comparisons += len(A) - 1
	if len(A) < 2 {
		return A
	}
	pivot := selectPivot(A)

	// partition Array
	i := pivot + 1
	for j := pivot + 1; j < len(A); j++ {
		if A[j] < A[pivot] {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[i-1], A[pivot] = A[pivot], A[i-1]

	// call partition on left and right sides of array
	partitionArray(A[:i-1])
	partitionArray(A[i:])

	return A
}

func selectPivot(A []int) int {
	return 0
}
