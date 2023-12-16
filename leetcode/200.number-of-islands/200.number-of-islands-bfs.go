//BFS

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	landCount := 0
	queue := [][2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				landCount = landCount + 1
				grid[i][j] = '0'
				queue = append(queue, [2]int{i, j})
				offsets := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
				for len(queue) > 0 {
					currPair := queue[0]
					queue = queue[1:]
					row := len(grid)
					col := len(grid[0])
					for _, offset := range offsets {
						m := currPair[0] + offset[0]
						n := currPair[1] + offset[1]
						if m >= 0 && m < row && n >= 0 && n < col && grid[m][n] == '1' {
							grid[m][n] = '0'
							queue = append(queue, [2]int{m, n})
						}
					}
				}
			}
		}
	}
	return landCount
}
