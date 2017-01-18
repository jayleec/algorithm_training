package algoutils

func Swap(A []int, a, b int){
	tmp := A[a]
	A[a] = A[b]
	A[b] = tmp
}
