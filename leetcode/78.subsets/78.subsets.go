package main

import "fmt"

type subset [][]int

func (s *subset) backtrack(path []int, nums []int, start int) {
	//fmt.Printf("%vbacktrack(path: %v, nums: %v, start: %v, originalStart: %v) \n", getIndentation(originalStart), path, nums, start, originalStart)

	tmp := make([]int, len(path))
	copy(tmp, path)
	*s = append(*s, tmp)

	for i := start; i < len(nums); i++ {
		//fmt.Printf("Added element %v \n", nums[i])
		path = append(path, nums[i])
		s.backtrack(path, nums, i+1)
		//fmt.Printf("removing element %v \n", path[len(path)-1])
		path = path[:len(path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	s := subset{}
	s.backtrack([]int{}, nums, 0)

	fmt.Println(s)
}
