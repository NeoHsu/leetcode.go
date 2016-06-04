package addDigits

import (
	"testing"
)

func Test_addDigits_1(t *testing.T) {
	result := addDigits(38)
	if result == 2 {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %d", result)
	}
}

func Test_addDigits_2(t *testing.T) {
	result := addDigits(11)
	if result == 2 {
		t.Log("Test 2 pass.")
	} else {
		t.Errorf("Test 2 fail. result: %d", result)
	}
}

func Test_addDigits_3(t *testing.T) {
	result := addDigits(0)
	if result == 0 {
		t.Log("Test 3 pass.")
	} else {
		t.Errorf("Test 3 fail. result: %d", result)
	}
}
