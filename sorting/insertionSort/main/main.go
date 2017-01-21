package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main(){

	slice := SliceGenerator(10)
	fmt.Println("Unsorted: ", slice)

	fmt.Println("Sorted: ", InsertionSort(slice))
}

func InsertionSort(slice []int) []int{
	v, j := 0,0
	for i:=1 ; i< len(slice); i++{
		v = slice[i]
		j = i
		for j >= 1 && slice[j - 1] > v{
			slice[j] = slice[j - 1]
			j = j - 1
		}
		slice[j] = v
	}
	return slice
}

func SliceGenerator(size int)[]int {
	rand.Seed(time.Now().UnixNano())
	temp := []int{}
	for i := 0 ; i< size ; i++{
		temp = append(temp, rand.Intn(88))
	}
	return temp
}
