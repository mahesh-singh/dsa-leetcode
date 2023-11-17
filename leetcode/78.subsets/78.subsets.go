package main

import "fmt"

type subset [][]int

func (s *subset) backtrack(path []int, nums []int, start int, whoiscalling string) {
	fmt.Printf("backtrack(path: %v, nums: %v, start: %v, whoiscalling : %v) \n", path, nums, start, whoiscalling)

	if start >= len(nums) {
		fmt.Printf("adding into result: %v \n", path)
		tmp := make([]int, len(path))
		copy(tmp, path)
		*s = append(*s, tmp)
		return
	}

	fmt.Printf("append into path: %v, start:%v \n", path, start)
	path = append(path, nums[start])
	s.backtrack(path, nums, start+1, "adding")

	fmt.Printf("remove into path: %v start:%v \n", path, start)
	path = path[:len(path)-1]
	s.backtrack(path, nums, start+1, "removing")
	/*
		for i := start; i < len(nums); i++ {
			fmt.Printf("Added element %v \n", nums[i])
			path = append(path, nums[i])
			fmt.Printf("path %v \n", path)
			s.backtrack(path, nums, i+1)
			fmt.Printf("removing element %v \n", path[len(path)-1])
			path = path[:len(path)-1]
			fmt.Printf("path>> %v \n", path)
		}
	*/

	fmt.Printf("func end...... start:%v \n", start)
}

func main() {
	nums := []int{1, 2, 3}
	s := subset{}
	s.backtrack([]int{}, nums, 0, "root")

	fmt.Println(s)
}
