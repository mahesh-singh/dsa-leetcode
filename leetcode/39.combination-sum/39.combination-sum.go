package main

import "fmt"

type result [][]int

func combinationSum(candidates []int, target int) [][]int {
	res := result{}
	res.dfs(candidates, target, []int{}, 0, 0)
	return res
}

func (r *result) dfs(candidate []int, target int, path []int, sum int, start int) {

	if sum > target || start >= len(candidate) {
		return
	}
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*r = append(*r, tmp)
		return
	}
	path = append(path, candidate[start])
	sum = sum + candidate[start]
	r.dfs(candidate, target, path, sum, start)

	path = path[:len(path)-1]
	sum = sum - candidate[start]
	r.dfs(candidate, target, path, sum, start+1)

}

func main() {
	r := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(r)
}
