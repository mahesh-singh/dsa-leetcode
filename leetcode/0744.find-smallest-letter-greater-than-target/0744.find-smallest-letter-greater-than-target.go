package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {

	l := 0
	r := len(letters) - 1
	ret := letters[0]

	for l <= r {
		mid := l + (r-l)/2

		if letters[mid] == target {
			if mid > len(letters) { // Missing edge case
				return letters[mid+1]
			}
		}

		if letters[mid] > target {
			ret = letters[mid]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ret

}
