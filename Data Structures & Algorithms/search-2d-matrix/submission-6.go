func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)*len(matrix[0])-1
	for left <= right {
		mid := left + (right-left)/2
		val := matrix[mid/len(matrix[0])][mid%len(matrix[0])]
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
