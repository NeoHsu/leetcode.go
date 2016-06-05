package removeDuplicatesFromSortedArrayII

import (
	"testing"
)

func Test_removeDuplicates_1(t *testing.T) {
	result := removeDuplicates([]int{1, 1, 1, 2, 2, 3})
	if result == 5 {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %d", result)
	}
}

func Test_removeDuplicates_2(t *testing.T) {
	result := removeDuplicates([]int{1, 1, 1, 2, 2, 3, 4, 4, 4, 5, 5})
	if result == 9 {
		t.Log("Test 2 pass.")
	} else {
		t.Errorf("Test 2 fail. result: %d", result)
	}
}
