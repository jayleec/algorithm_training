package algoutils

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int)[]int {
	slice :=[]int{}
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<size; i++{
		slice = append(slice, rand.Intn(80))
	}
	return slice
}
