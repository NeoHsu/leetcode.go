package reverseString

// 透過交換的方式去做 reverse

func reverseString(s string) (str string) {
	r := []int32(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	str = string(r)
	return
}
