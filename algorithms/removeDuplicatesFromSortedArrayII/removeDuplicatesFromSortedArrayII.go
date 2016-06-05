package removeDuplicatesFromSortedArrayII

func removeDuplicates(nums []int) int {
	current := 0
	lastElement := 0
	count := 0
	for i, v := range nums {
		if i == 0 {
			current = 0
			lastElement = v
			count = 1
		} else {
			if lastElement == v {
				count++
				if count <= 2 {
					current++
					nums[current] = v
				}
			} else {
				lastElement = v
				count = 1
				current++
				nums[current] = v
			}
		}
	}
	if len(nums) == 0 {
		return len(nums[:])
	} else {
		return len(nums[:current+1])
	}

}
