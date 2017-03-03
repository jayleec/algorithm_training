package main

import "fmt"

func main(){
	n := 10
	memo = make([]int, n)
	fmt.Println(fib(n))
}

var memo []int

func fib(n int) int {
	var f int
	if memo[n-1] != 0{
		return memo[n-1]
	}else if n <= 2 {
		f = 1
	}else {
		f = fib(n - 1) + fib(n - 2)
		memo[n-1] = f
	}
	return f
}