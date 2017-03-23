package main

import (
	"github.com/jayleec/algorithm_training/algoutils"
	"fmt"
)

func main(){
	testCase := algoutils.GenerateSlice(10)
	fmt.Println(testCase)
	fmt.Println(MergeSort(testCase))
}

func MergeSort(A []int) []int{
	left, size := 0,0
	for size = 1; size <=len(A)-1; size += size {
		for left=0; left<len(A)-2; left += size*2 {
			mid := left+size-1
			rightEnd := Min(left+2*size-1, len(A)-1)
			Merge(A, left, mid, rightEnd)
		}
	}
	return A
}

func Merge(A []int, l int, m int, r int) {
	n1 := m-l+1
	n2 := r-m
	//create buffer arrays
	L := make([]int, n1)
	R := make([]int, n2)

	//copy data to temp arrays
	for i:=0; i<n1; i++{
		L[i] = A[l+i]
	}
	for j:=0; j<n2; j++{
		R[j] = A[m+1+j]
	}
	i, j, k := 0,0,l
	for i < n1 && j < n2  {
		if L[i] <= R[j]{
			A[k] = L[i]
			i = i+1
		}else{
			fmt.Println("k", k)
			A[k] = R[j]
			j = j+1
		}
		k = k+1
	}

	for i< n1{
		A[k] = L[i]
		i = i+1
		k = k +1
	}
	for j <n2{
		A[k] = R[j]
		j = j+1
		k = k+1
	}
}

func Min(a, b int) int{
	if a < b{
		return a
	}
	return b
}


