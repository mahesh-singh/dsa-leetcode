func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				islandArea := dfs(&grid, [2]int{i, j})
				if islandArea > maxArea {
					maxArea = islandArea
				}
			}
		}
	}
	return maxArea
}

func dfs(grid *[][]int, curr [2]int) int {
	area := 1
	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, offset := range offsets {
		x := curr[0] + offset[0]
		y := curr[1] + offset[1]
		row := len((*grid))
		col := len((*grid)[0])
		if x >= 0 && x < row && y >= 0 && y < col && (*grid)[x][y] == 1 {
			(*grid)[x][y] = 0
			area = area + dfs(grid, [2]int{x, y})
		}
	}
	return area
}