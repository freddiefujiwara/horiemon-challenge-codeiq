package Horiemon

func Target(input []int) int {
	return input[0]
}
func Except(pair []int, array []int) []int {
	var except []int
	var start int
	for i, v := range array {
		if v == pair[1] {
			start = i + 1
		}

	}
	for _, v := range array[start:] {
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
func Calculate(target int, array []int) [][]int {
	var answeres [][]int
	for _, pair := range ListUp(array) {
		for _, v := range Except(pair, array) {
			if target == pair[0]+pair[1]+v {
				answeres = append(answeres, []int{pair[0], pair[1], v})
			}
		}
	}
	return answeres
}
