package Horiemon

import (
	"testing"
)

func TestSum(t *testing.T) {
	actual := Target([]int{1, 2, 3, 4, 5})
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
