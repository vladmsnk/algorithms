package tasks

import "sort"

func groupAnagrams(strs []string) [][]string {

	var res [][]string
	mp := make(map[string][]string)

	for _, str := range strs {
		s := []rune(str)
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		key := string(s)
		mp[key] = append(mp[key], str)
	}
	for _, val := range mp {
		res = append(res, val)
	}
	return res
}
