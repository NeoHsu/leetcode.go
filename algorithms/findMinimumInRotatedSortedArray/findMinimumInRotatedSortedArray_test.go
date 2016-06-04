package findMinimumInRotatedSortedArray

import (
	"testing"
)

func Test_findMin_1(t *testing.T) {
	result := findMin([]int{0, 1, 2, 4, 5, 6, 7})
	if result == 0 {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %d", result)
	}
}

func Test_findMin_2(t *testing.T) {
	result := findMin([]int{6, 1, 9, 4, 5, 0})
	if result == 0 {
		t.Log("Test 2 pass.")
	} else {
		t.Errorf("Test 2 fail. result: %d", result)
	}
}

func Test_findMin_3(t *testing.T) {
	result := findMin([]int{6, 9, 9, 4, 5, 9, 4, 5, 7})
	if result == 4 {
		t.Log("Test 3 pass.")
	} else {
		t.Errorf("Test 3 fail. result: %d", result)
	}
}
