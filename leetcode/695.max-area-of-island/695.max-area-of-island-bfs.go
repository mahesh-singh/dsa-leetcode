func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	queue := [][2]int{}
	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	row := len(grid)
	col := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				islandArea := 1
				queue = append(queue, [2]int{i, j})
				grid[i][j] = 0
				for len(queue) > 0 {
					currPair := queue[0]
					queue = queue[1:]
					for _, offset := range offsets {
						x := currPair[0] + offset[0]
						y := currPair[1] + offset[1]
						if x >= 0 && x < row && y >= 0 && y < col && grid[x][y] == 1 {
							queue = append(queue, [2]int{x, y})
							grid[x][y] = 0
							islandArea = islandArea + 1
						}
					}
					if islandArea > maxArea {
						maxArea = islandArea
					}
				}
			}
		}
	}
	return maxArea
}