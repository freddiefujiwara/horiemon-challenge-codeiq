package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFact(t *testing.T) {
	if Fact(0) != 0 {
		t.Errorf("error")
	}

	if Fact(1) != 1 {
		t.Errorf("error")
	}

	if Fact(2) != 3 {
		t.Errorf("error")
	}

	if Fact(3) != 6 {
		t.Errorf("error")
	}

	if Fact(4) != 10 {
		t.Errorf("error")
	}

	if Fact(5) != 15 {
		t.Errorf("error")
	}

	if Fact(6) != 21 {
		t.Errorf("error")
	}

	if Fact(7) != 28 {
		t.Errorf("error")
	}

	if Fact(8) != 36 {
		t.Errorf("error")
	}

	if Fact(9) != 45 {
		t.Errorf("error")
	}
}

func TestListUp(t *testing.T) {
	array := []int{8, 4, 10, 3, 2}
	sort.Ints(array)
	actual := ListUp(15, 5, array)

	expected := [][]int{{2, 3, 1}, {2, 4, 2}, {2, 8, 3}, {2, 10, 4}, {3, 4, 2}, {3, 8, 3}, {3, 10, 4}, {4, 8, 3}, {4, 10, 4}, {}}
	if !reflect.DeepEqual(len(actual), len(expected)) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestExcept(t *testing.T) {
	actual := Except([]int{5, 8, 1}, 5, []int{5, 8, 4, 10, 3, 2})
	expected := []int{4, 10, 3, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	actual = Except([]int{8, 4, 2}, 5, []int{5, 8, 4, 10, 3, 2})
	expected = []int{10, 3, 2}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCalculate(t *testing.T) {
	array := []int{8, 4, 10, 3, 2}
	sort.Ints(array)
	actual := Calculate(15, 5, array)
	expected := [][]int{{2, 3, 10}, {3, 4, 8}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
	array = []int{13, 12, 17, 10, 4, 18, 3, 11, 5, 7}
	sort.Ints(array)
	actual = Calculate(35, 10, array)
	expected = [][]int{{4, 13, 18}, {5, 12, 18}, {5, 13, 17}, {7, 10, 18}, {7, 11, 17}, {10, 12, 13}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
