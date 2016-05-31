package longestCommonPrefix

import (
	"testing"
)

func Test_longestCommonPrefix_1(t *testing.T) {
	data := []string{"a", "apple", "add", "animal"}
	result := longestCommonPrefix(data)
	if result != "a" {
		t.Errorf("Test 1 fail. result: %s", result)
	} else {
		t.Log("Test 1 pass.")
	}
}

func Test_longestCommonPrefix_2(t *testing.T) {
	data := []string{"apple", "a", "cat", "animal"}
	result := longestCommonPrefix(data)
	if result != "" {
		t.Errorf("Test 2 fail. result: %s", result)
	} else {
		t.Log("Test 2 pass.")
	}
}

func Test_longestCommonPrefix_3(t *testing.T) {
	data := []string{}
	result := longestCommonPrefix(data)
	if result != "" {
		t.Errorf("Test 3 fail. result: %s", result)
	} else {
		t.Log("Test 3 pass.")
	}
}

func Test_longestCommonPrefix_4(t *testing.T) {
	data := []string{"apple", "apple"}
	result := longestCommonPrefix(data)
	if result != "apple" {
		t.Errorf("Test 4 fail. result: %s", result)
	} else {
		t.Log("Test 4 pass.")
	}
}

func Test_longestCommonPrefix_5(t *testing.T) {
	data := []string{"cat"}
	result := longestCommonPrefix(data)
	if result != "cat" {
		t.Errorf("Test 5 fail. result: %s", result)
	} else {
		t.Log("Test 5 pass.")
	}
}
