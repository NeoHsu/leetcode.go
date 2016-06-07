package singleNumberIII

import (
	"testing"
)

func Test_singleNumber_1(t *testing.T) {
	result := singleNumber([]int{1, 2, 1, 3, 2, 5})
	if result[0] == 3 && result[1] == 5 {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %d .", result)
	}
}
