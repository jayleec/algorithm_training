package algoutils


func InsertSlice(slice, new []int, index int) []int {
	return append(slice[:index], append(new, slice[index:]...)...)
}

func Prepend(slice []int, elem int) []int {
	temp := []int{elem}
	for _, v :=  range slice {
		temp = append(temp, v)
	}
	return temp
}