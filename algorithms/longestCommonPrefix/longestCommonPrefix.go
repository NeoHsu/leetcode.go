package longestCommonPrefix

// 取第一個 array element 一個個字元和其他 element 做比對

func longestCommonPrefix(strs []string) (str string) {
	if len(strs) == 0 {
		return
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				str = strs[0][:i]
				return
			}
		}
	}
	str = strs[0]
	return
}
