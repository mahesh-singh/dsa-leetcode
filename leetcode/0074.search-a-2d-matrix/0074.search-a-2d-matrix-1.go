package leetcode

func searchMatrix(matrix [][]int, target int) bool {

	//Mistake: as i is increasing and j is decreasing, only need to check boundary condition i < len(matrix) && j >=0
	for i, j := 0, len(matrix[0])-1; i < len(matrix) && j >= 0; {
		if matrix[i][j] == target {
			return true
		}

		//Mistake: increasing i instead of decreasing j
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
