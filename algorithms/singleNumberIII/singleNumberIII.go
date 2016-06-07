package singleNumberIII

func singleNumber(nums []int) []int {
	t := make(map[int]int)
	r := []int{}
	for _, v := range nums {
		t[v]++
	}
	for k, v := range t {
		if v == 1 {
			r = append(r, k)
		}
	}
	return r
}
