package searchA2DMatrix

func searchMatrix(matrix [][]int, target int) (isHit bool) {
	for _, v := range matrix {
		low, height := 0, len(v)-1
		for i := 0; i < len(v); i++ {
			if low > height {
				break
			}
			if v[low] == target || v[height] == target {
				isHit = true
				return
			}
			low++
			height--
		}
	}
	return
}
