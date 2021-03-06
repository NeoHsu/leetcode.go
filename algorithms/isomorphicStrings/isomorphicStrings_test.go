package isomorphicStrings

import (
	"testing"
)

func Test_isIsomorphic_1(t *testing.T) {
	result := isIsomorphic("egg", "add")
	if result == true {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %t .", result)
	}
}

func Test_isIsomorphic_2(t *testing.T) {
	result := isIsomorphic("foo", "bar")
	if result == false {
		t.Log("Test 2 pass.")
	} else {
		t.Errorf("Test 2 fail. result: %t .", result)
	}
}

func Test_isIsomorphic_3(t *testing.T) {
	result := isIsomorphic("paper", "title")
	if result == true {
		t.Log("Test 3 pass.")
	} else {
		t.Errorf("Test 3 fail. result: %t .", result)
	}
}

func Test_isIsomorphic_4(t *testing.T) {
	result := isIsomorphic("", "")
	if result == true {
		t.Log("Test 4 pass.")
	} else {
		t.Errorf("Test 4 fail. result: %t .", result)
	}
}

/*
func Test_twoSum_2(t *testing.T) {
	data := []int{2, 7, 11, 15}
	result := twoSum(data, 100)
	if result[0] != 0 && result[1] != 0 {
		t.Errorf("Test 2 fail. result[0]: %d result[1]: %d .", result[0], result[1])
	} else {
		t.Log("Test 2 pass.")
	}
}

func Test_twoSum_3(t *testing.T) {
	data := []int{2, -7, 11, 15, -30, 22, 4, -6, 33}
	result := twoSum(data, 3)
	if result[0] != 4 && result[1] != 8 {
		t.Errorf("Test 3 fail. result[0]: %d result[1]: %d .", result[0], result[1])
	} else {
		t.Log("Test 3 pass.")
	}
}

func Test_twoSum_4(t *testing.T) {
	data := []int{2, -7, 11, 15, 30, 22, 4, -6, -33}
	result := twoSum(data, -3)
	if result[0] != 1 && result[1] != 6 {
		t.Errorf("Test 4 fail. result[0]: %d result[1]: %d .", result[0], result[1])
	} else {
		t.Log("Test 4 pass.")
	}
}
*/
