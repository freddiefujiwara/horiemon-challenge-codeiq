package main

import (
	"reflect"
	"testing"
)

func TestListUp(t *testing.T) {
	actual := ListUp([]int{8, 4, 10, 3, 2})
	expected := [][]int{
		{8, 4}, {8, 10}, {8, 3}, {8, 2},
		{4, 10}, {4, 3}, {4, 2},
		{10, 3}, {10, 2},
		{3, 2}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestExcept(t *testing.T) {
	actual := Except([]int{5, 8}, []int{5, 8, 4, 10, 3, 2})
	expected := []int{4, 10, 3, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual = Except([]int{8, 4}, []int{5, 8, 4, 10, 3, 2})
	expected = []int{10, 3, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestCalculate(t *testing.T) {
	actual := Calculate(15, []int{8, 4, 10, 3, 2})
	expected := [][]int{{8, 4, 3}, {10, 3, 2}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual = Calculate(35, []int{13, 12, 17, 10, 4, 18, 3, 11, 5, 7})
	expected = [][]int{{13, 12, 10}, {13, 17, 5}, {13, 4, 18}, {12, 18, 5}, {17, 11, 7}, {10, 18, 7}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
