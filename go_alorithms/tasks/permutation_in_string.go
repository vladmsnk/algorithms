package tasks

/*
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
In other words, return true if one of s1's permutations is the substring of s2.
*/

//Ex1 s1 = "ab", s2 = "eidbaooo" ->true
//Ex2 s1 = "ab", s2 = "eidboaoo" ->false

func checkInclusion(s1 string, s2 string) bool {
	countMap := make(map[rune]int)

	for _, el := range s1 {
		countMap[el]++
	}

	for l, r := 0, 0; r < len(s2); {

		currElem := rune(s2[r])

		cnt := countMap[currElem]
		if cnt == 0 {
			if l != r {
				countMap[rune(s2[l])]++
			} else {
				r++
			}
			l++
		} else {
			countMap[currElem]--
			if r-l+1 == len(s1) {
				return true
			}
			r++
		}
	}

	return false
}
