package main

import (
	"fmt"
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
		} else {
			if len(array) < count {
				array = append(array, v)
			}
		}
		input_count++
	}
	fmt.Println(len(Calculate(target, array)))
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
