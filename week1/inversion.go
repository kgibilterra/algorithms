package mergesort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inversions int

// Inversions gets the data and starts the inversion divide step
func Inversions(data []int) {
	finalArr := InversionDivide(openData())
	fmt.Println("inversions: ", inversions, "length of final data: ", len(finalArr))
}

// InversionDivide divides up the array and recurses until all inversion
// are counted
func InversionDivide(a []int) []int {
	if len(a) == 1 {
		return a
	}
	return Combine(InversionDivide(a[0:len(a)/2]), InversionDivide(a[len(a)/2:]))
}

// Combine returns a sorted array from two previously sorted arrays
func Combine(b, c []int) []int {
	var r []int
	i, j := 0, 0

	for {
		// check if out of bounds
		if j == len(c) {
			r = append(r, b[i:]...)
			return r
		} else if i == len(b) {
			r = append(r, c[j:]...)
			return r
		}

		// find which is less
		if b[i] < c[j] {
			r = append(r, b[i])
			i++
		} else {
			r = append(r, c[j])
			inversions += len(b) - i
			j++
		}
	}
}

func openData() []int {
	f, err := os.Open("data.txt")
	if err != nil {
		panic("error reading file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	data := []int{}
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		data = append(data, x)
	}
	return data
}
