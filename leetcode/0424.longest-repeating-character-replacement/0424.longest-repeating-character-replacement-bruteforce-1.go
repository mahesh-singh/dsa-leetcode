package leetcode

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	frqMap := make(map[byte]int)
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if frqVal, ok := frqMap[s[j]]; ok {
				frqMap[s[j]] = frqVal + 1
			} else {
				frqMap[s[j]] = 1
			}
			frqCharCount := mostFrqCharCount(frqMap)
			currLen := j - i + 1
			if currLen-frqCharCount <= k {
				if currLen > maxLen {
					maxLen = currLen
				}
			} else {
				frqMap = make(map[byte]int)
				break
			}
		}
	}
	return maxLen
}

func mostFrqCharCount(m map[byte]int) int {
	maxCount := 0
	for _, frq := range m {
		if frq > maxCount {
			maxCount = frq
		}
	}
	return maxCount
}
