package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	frqMap := make(map[byte]int)
	maxLen := 1
	l := 0
	r := 0
	for r < len(s) {
		if pos, ok := frqMap[s[r]]; ok && pos >= l {
			l = pos + 1
		}
		frqMap[s[r]] = r
		length := r - l + 1
		if length > maxLen {
			maxLen = length
		}
		r = r + 1
	}

	return maxLen
}
