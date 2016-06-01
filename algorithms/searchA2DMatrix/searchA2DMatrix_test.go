package searchA2DMatrix

import (
	"testing"
)

func Test_searchMatrix_1(t *testing.T) {
	data := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	result := searchMatrix(data, 3)
	if result == true {
		t.Log("Test 1 pass.")
	} else {
		t.Errorf("Test 1 fail. result: %t ", result)
	}
}

func Test_searchMatrix_2(t *testing.T) {
	data := [][]int{{1, 3, 8, 5, 7}, {10, 11, 90, 16, 20}, {23, 30, 34, 50, 22}}
	result := searchMatrix(data, 90)
	if result == true {
		t.Log("Test 2 pass.")
	} else {
		t.Errorf("Test 2 fail. result: %t ", result)
	}
}

func Test_searchMatrix_3(t *testing.T) {
	data := [][]int{{1, 3, 7}, {10, 11, 90, 16, 20}, {23, 30, 34, 50, 30, 34, 50, 22}}
	result := searchMatrix(data, 11)
	if result == true {
		t.Log("Test 3 pass.")
	} else {
		t.Errorf("Test 3 fail. result: %t ", result)
	}
}

func Test_searchMatrix_4(t *testing.T) {
	data := [][]int{{1}}
	result := searchMatrix(data, 1)
	if result == true {
		t.Log("Test 4 pass.")
	} else {
		t.Errorf("Test 4 fail. result: %t ", result)
	}
}

func Test_searchMatrix_5(t *testing.T) {
	data := [][]int{{1}}
	result := searchMatrix(data, 2)
	if result == false {
		t.Log("Test 5 pass.")
	} else {
		t.Errorf("Test 5 fail. result: %t ", result)
	}
}

func Test_searchMatrix_6(t *testing.T) {
	data := [][]int{}
	result := searchMatrix(data, 2)
	if result == false {
		t.Log("Test 6 pass.")
	} else {
		t.Errorf("Test 6 fail. result: %t ", result)
	}
}
