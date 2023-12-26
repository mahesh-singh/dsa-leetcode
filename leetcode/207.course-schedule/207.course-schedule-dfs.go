func canFinish(numCourses int, prerequisites [][]int) bool {
	adjMap := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		course := prerequisites[i][0]
		pre := prerequisites[i][1]
		adjMap[pre] = append(adjMap[pre], course)
	}
	visited := make([]bool, numCourses)
	visiting := make([]bool, numCourses)

	for key := range adjMap {
		if !dfs(key, &visiting, &visited, &adjMap) {
			return false
		}
	}
	return true
}

func dfs(curr int, visiting *[]bool, visited *[]bool, adjMap *map[int][]int) bool {
	if (*visiting)[curr] {
		return false
	}
	if (*visited)[curr] {
		return true
	}
	(*visiting)[curr] = true
	adj := (*adjMap)[curr]

	for i := 0; i < len(adj); i++ {
		if (*visited)[adj[i]] == false {
			if !dfs(adj[i], visiting, visited, adjMap) {
				return false
			}
		}
	}
	(*visiting)[curr] = false
	(*visited)[curr] = true
	return true

}