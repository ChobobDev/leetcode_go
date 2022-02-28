package easy

// Runtime: 3 ms, faster than 42.06% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.3 MB, less than 75.72% of Go online submissions for Longest Common Prefix.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var CommonPrefix string
	for i := 0; ; i++ {
		for j := 0; j <= len(strs)-2; j++ {
			if len(strs[j]) < i || len(strs[j+1]) < i || strs[j][:i] != strs[j+1][:i] {
				return CommonPrefix
			}
		}
		CommonPrefix = strs[0][:i]
	}
	return CommonPrefix
}
