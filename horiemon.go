package Horiemon

func Target(input []int) int {
	return input[0]
}
func Except(pair []int, array []int) []int {
	var except []int
	for _, v := range array {
		if v != pair[0] && v != pair[1] {
			except = append(except, v)
		}
	}
	return except
}

func ListUp(array []int) [][]int {
	var list [][]int
	for i, head := range array {
		for _, v := range array[i+1:] {
			list = append(list, []int{head, v})
		}
	}
	return list
}
