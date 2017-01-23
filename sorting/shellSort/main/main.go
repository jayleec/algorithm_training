package main

import (
	"fmt"
	"github.com/jayleec/algorithm_training/algoutils"
)

func main() {

	slice := algoutils.GenerateSlice(10)
	fmt.Println("Unsorted: ", slice)
	fmt.Println("Sorted: ", shellSort(slice))

}

func shellSort(slice []int) []int{
	var inner, outer, interval, temp int
	interval = 1
	for interval < len(slice){
		interval = interval * 3 + 1
	}

	for interval > 0{
		for outer = interval; outer < len(slice) ;outer++{
			temp = slice[outer]
			inner = outer
			for inner >= interval && slice[inner - interval] > temp {
				slice[inner] = slice[inner - interval]
				inner = inner - interval
			}
			slice[inner] = temp
		}
		interval = interval / 3
	}

	return slice
}
