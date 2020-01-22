//Develop a program which solves the following task:
//Find one of the numbers which exists in each of three nondecreasing arrays x[p], y[q], z[r]. Algorithm
//complexity should be O(p+q+r).

package problems

import "fmt"

func RunFindCommonNumberProblem() {
	x := []int{1, 5, 10, 20, 40, 80}
	y := []int{6, 7, 20, 80, 100}
	z := []int{3, 4, 15, 20, 30, 70, 80, 120}

	result := findCommonNumber3(x, y, z)

	fmt.Printf("common number is: %d\n", result)
}

func findCommonNumber(x, y, z []int) int {
	xIt, yIt, zIt := 0, 0, 0

	var result int = -1

	for ; xIt < len(x); xIt++ {
		for ; yIt < len(y) && y[yIt] < x[xIt]; yIt++ {
		}
		for ; zIt < len(z) && z[zIt] < x[xIt]; zIt++ {
		}

		if x[xIt] == y[yIt] && x[xIt] == z[zIt] {
			result = x[xIt]
			break
		}
	}

	return result
}

func findCommonNumber2(x, y, z []int) int {
	xIt, yIt, zIt := 0, 0, 0

	var result int = -1

OUTER:
	for xIt < len(x) && yIt < len(y) {
		if x[xIt] < y[yIt] {
			xIt++
		} else if x[xIt] > y[yIt] {
			yIt++
		} else {
			for zIt < len(z) {
				if x[xIt] > z[zIt] {
					zIt++
				} else if x[xIt] == z[zIt] {
					result = x[xIt]
					break OUTER
				}
			}
		}
	}

	return result
}

func findCommonNumber3(x, y, z []int) int {
	xIt, yIt, zIt := 0, 0, 0

	var result int = -1

	for xIt < len(x) && yIt < len(y) && zIt < len(z) {
		if x[xIt] == y[yIt] && y[yIt] == z[zIt] {
			result = x[xIt]
			break
		} else if x[xIt] < y[yIt] {
			xIt++
		} else if y[yIt] < z[zIt] {
			yIt++
		} else {
			zIt++
		}
	}

	return result
}
