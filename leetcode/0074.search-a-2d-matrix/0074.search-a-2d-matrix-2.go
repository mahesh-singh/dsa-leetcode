package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	top := 0
	bot := len(matrix) - 1
	row := 0

	// find candidate row
	for top <= bot {
		row = (top + bot) / 2

		if matrix[row][0] > target {
			bot = row - 1
		} else if matrix[row][len(matrix[row])-1] < target {
			top = row + 1
		} else {
			// we are on the curren row
			break
		}
	}

	if top > bot {
		return false
	}

	l := 0
	r := len(matrix[row]) - 1

	for l <= r {
		mid := l + (r-l)/2

		if matrix[row][mid] == target {
			return true
		}

		if matrix[row][mid] > target {
			r = mid - 1
		} else if matrix[row][mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
