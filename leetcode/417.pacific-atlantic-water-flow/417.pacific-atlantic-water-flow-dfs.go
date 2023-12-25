func pacificAtlantic(heights [][]int) [][]int {
	ret := [][]int{}
	row := len(heights)
	col := len(heights[0])

	pacific := make([][]bool, row)
	atlantic := make([][]bool, row)

	for r := 0; r < row; r++ {
		pacific[r] = make([]bool, col)
		atlantic[r] = make([]bool, col)
	}

	for c := 0; c < col; c++ {
		dfs(heights, [2]int{0, c}, [2]int{0, c}, &pacific)
		dfs(heights, [2]int{row - 1, c}, [2]int{row - 1, c}, &atlantic)
	}
	for r := 0; r < row; r++ {
		dfs(heights, [2]int{r, 0}, [2]int{r, 0}, &pacific)
		dfs(heights, [2]int{r, col - 1}, [2]int{r, col - 1}, &atlantic)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return ret
}

func dfs(heights [][]int, currPair [2]int, prePair [2]int, visited *[][]bool) {
	r := currPair[0]
	c := currPair[1]
	row := len(heights)
	col := len(heights[0])

	if r < 0 || r == row || c < 0 || c == col || (*visited)[r][c] == true || heights[prePair[0]][prePair[1]] > heights[r][c] {
		return
	}

	(*visited)[r][c] = true
	dfs(heights, [2]int{r - 1, c}, [2]int{r, c}, visited)
	dfs(heights, [2]int{r + 1, c}, [2]int{r, c}, visited)
	dfs(heights, [2]int{r, c - 1}, [2]int{r, c}, visited)
	dfs(heights, [2]int{r, c + 1}, [2]int{r, c}, visited)
	//offsets := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	/*
			for _, offset := range offsets {
				x := r + offset[0]
				y := c + offset[1]

		        if x < 0 || x == row || y < 0 || y == col || (*visited)[x][y] == true || heights[r][c] > heights[x][y] {
		            return
		        }

				(*visited)[x][y] = true
		        dfs(heights, [2]int{x, y}, visited)
			}*/

}