package main

func findKthLargest(nums []int, k int) int {
	k = len(nums) - k
	return qs(nums, k)
}

func qs(nums []int, k int) int {
	start, end := 0, len(nums)-1
	p := partition(nums, start, end)

	if k == p {
		return nums[p]
	} else if k < p {
		return qs(nums[0:p], k)
	} else {
		return qs(nums[p+1:], k-(p+1))
	}
}

func partition(nums []int, start int, end int) int {

	pivot := nums[end]

	i := start

	for j := 0; j < end; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i = i + 1
		}

	}

	nums[i], nums[end] = nums[end], nums[i]
	return i
}
