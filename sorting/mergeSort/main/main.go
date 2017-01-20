package main

import "fmt"

func main(){
	tempList := []int{6,5,7,4,8,3,2,9,10,1}
	fmt.Println(MergeSort(tempList))

}

func MergeSort(slice []int) []int{
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func Merge(left, right []int) []int {
	size, i, j := len(left) + len(right), 0, 0
	slice := make([]int, size)

	for k := 0 ; k <size; k++{
		if i > len(left) - 1 && j<= len(right) - 1{
			slice[k] = right[j]
			j = j + 1
		}else if j > len(right) - 1 && i <= len(left) - 1{
			slice[k] = left[i]
			i = i + 1
		}else if left[i] < right[j]{
			slice[k] = left[i]
			i = i + 1
		}else{
			slice[k] = right[j]
			j = j + 1
		}
	}
	return slice
}



