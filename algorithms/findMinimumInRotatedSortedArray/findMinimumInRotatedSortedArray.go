package findMinimumInRotatedSortedArray

func findMin(nums []int) int {
	for _, v := range nums {
		if v < nums[0] {
			nums[0] = v
		}
	}
	return nums[0]
}
