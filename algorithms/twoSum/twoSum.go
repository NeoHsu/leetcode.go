package twoSum

// 透過 Map 存取另一個數值當作 Key

func twoSum(nums []int, target int) []int {
	pairNum := make(map[int]int)

	for i, value := range nums {
		pairIndex, ok := pairNum[value]
		if ok {
			return []int{pairIndex, i}
		}
		pairNum[target-value] = i
	}

	return []int{0, 0}
}
