package leetcode

// Brute force
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	frqMap := make(map[byte]bool)
	maxLen := 1
	for i := 0; i < len(s); i++ {
		frqMap[s[i]] = true
		for j := i + 1; j < len(s); j++ {
			if _, ok := frqMap[s[j]]; ok {
				frqMap = make(map[byte]bool)
				break
			} else {
				frqMap[s[j]] = true
				if maxLen < len(frqMap) {
					maxLen = len(frqMap)
				}
			}
		}
	}

	return maxLen
}
