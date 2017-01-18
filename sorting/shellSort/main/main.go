package main

import (
	"fmt"
	"github.com/jayleec/algorithm_training/algoutils"
)

func main() {

	Array := []int{3, 7, 2, 8, 4, 1, 9, 10, 6, 5}

	shellSort(Array)

	fmt.Println(Array)
}

func shellSort(A []int) {
	h := 1
	for h < len(A) {
		h = h * 3 + 1
	}

	for h > 0 {
		for i := 1; i < len(A); i++{
			for j:=i ; j >= h && A[j- h] > A[j]; j = j - h {
				algoutils.Swap(A, j, j - h)
			}
		}

		h = h /3
	}
}
