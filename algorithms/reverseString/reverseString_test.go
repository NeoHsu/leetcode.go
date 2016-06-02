package reverseString

import (
	"testing"
)

func Test_reverseString_1(t *testing.T) {
	result := reverseString("Hello")
	if result == "olleH" {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %s", result)
	}
}

func Test_reverseString_2(t *testing.T) {
	result := reverseString("")
	if result == "" {
		t.Log("Test 2 pass.")
	} else {
		t.Errorf("Test 2 fail. result: %s", result)
	}
}

func Test_reverseString_3(t *testing.T) {
	result := reverseString("Hello World")
	if result == "dlroW olleH" {
		t.Log("Test 3 pass.")
	} else {
		t.Errorf("Test 3 fail. result: %s", result)
	}
}

func Test_reverseString_4(t *testing.T) {
	result := reverseString("Hello @ Google")
	if result == "elgooG @ olleH" {
		t.Log("Test 4 pass.")
	} else {
		t.Errorf("Test 4 fail. result: %s", result)
	}
}
