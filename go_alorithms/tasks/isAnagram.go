package tasks

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[rune]int)

	for _, el := range s {
		mp[el]++
	}
	for _, el := range t {
		res, ok := mp[el]
		if !ok {
			return false
		} else if ok && res == 1 {
			delete(mp, el)
		} else {
			mp[el]--
		}
	}
	if len(mp) == 0 {
		return true
	}
	return false
}
