//DFS

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	landCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				landCount = landCount + 1
				dfs(&grid, i, j)
			}
		}
	}
	return landCount
}

func dfs(grid *[][]byte, i int, j int) {
	m := len(*grid)
	n := len((*grid)[0])

	if i < m && j < n && i >= 0 && j >= 0 && (*grid)[i][j] == '1' {
		(*grid)[i][j] = '0' //visited
		dfs(grid, i-1, j)
		dfs(grid, i+1, j)
		dfs(grid, i, j-1)
		dfs(grid, i, j+1)
	} else {
		return
	}
}
