package Horiemon

import (
	"reflect"
	"testing"
)

func TestTarget(t *testing.T) {
	actual := Target([]int{15, 5, 8, 4, 10, 3, 2})
	expected := 15
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
func TestListUp(t *testing.T) {
	actual := ListUp([]int{5, 8, 4, 10, 3, 2})
	expected := [][]int{{5, 8}, {5, 4}, {5, 10}, {5, 3}, {5, 2},
		{8, 4}, {8, 10}, {8, 3}, {8, 2},
		{4, 10}, {4, 3}, {4, 2},
		{10, 3}, {10, 2},
		{3, 2}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
