package main

import (
	"fmt"
	"sort"
)

func main() {
	input_count := 0
	var target int
	var count int
	var v int
	var array []int
	for read_count, _ := fmt.Scan(&v); read_count != 0; read_count, _ = fmt.Scan(&v) {
		if 0 == input_count {
			target = v
		} else if 1 == input_count {
			count = v
			array = make([]int, count)
		} else {
			array[input_count-2] = v
		}
		input_count++
	}
	sort.Ints(array)
	fmt.Println(len(Calculate(target, count, array)))
}

func Except(pair []int, count int, array []int) []int {
	except := make([]int, count-pair[2])
	except_index := 0
	on := false
	for _, v := range array {
		if on {
			except[except_index] = v
			except_index++
		}
		if v == pair[1] {
			on = true
		}

	}
	return except
}

func ListUp(target int, count int, array []int) [][]int {
	list := make([][]int, Fact(count-1))
	list_index := 0
	for i, h := range array {
		if count > i+1 && target > h+array[i+1] {
			for j, v := range array[i+1:] {
				if count > i+j+1 && target > h+v {
					list[list_index] = []int{h, v, i + j + 1}
					list_index++
				}
			}
		}
	}
	return list
}

func Calculate(target int, count int, array []int) [][]int {
	var answeres [][]int
	for _, pair := range ListUp(target, count, array) {
		if 3 == len(pair) {
			pair_sum := pair[0] + pair[1]
			if target-pair_sum >= array[0] {
				for _, v := range Except(pair, count, array) {
					if target == pair[0]+pair[1]+v {
						answeres = append(answeres, []int{pair[0], pair[1], v})
					}
				}
			}
		}
	}
	return answeres
}

func Fact(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + Fact(n-1)
	}
}
