package tasks

// haystack = "leetcode", needle = "leeto"

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	l, r := 0, 0
	for r < len(haystack) && r-l != len(needle) {

		if haystack[r] != needle[r-l] {
			l++
			r = l
		} else {
			r++
		}
	}
	if r-l == len(needle) {
		return l
	}
	return -1
}
