func searchMatrix(matrix [][]int, target int) bool {

	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			// move to previous  column
			j = j - 1
		} else {
			// move to next row
			i = i + 1
		}

	}

	return false
}
