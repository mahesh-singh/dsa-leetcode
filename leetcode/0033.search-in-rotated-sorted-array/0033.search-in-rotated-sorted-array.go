func search(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			// increasing order on left side
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// increasing order side on right side
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

	}
	if nums[l] == target {
		return l
	} else {
		return -1
	}

}