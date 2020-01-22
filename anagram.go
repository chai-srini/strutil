package strutil

// IsAnagram checks weather two strings are anagrams of each other
// 1 - walk through s1 and populate hashtable with key = character and value = number of occurrences
// 2 - walk through s2 and check if char exists in hashtable. If does not exist, its not an anagram.
//  If exists, and value greater than 1, decrement it. If exists and value is 1, delete it.
// 3. - if any keys remaining in hashtable, it is not an anagram. Else it is an anagram
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return false
	}

	ht := make(map[rune]int)
	for _, c := range s1 {
		if _, ok := ht[c]; ok {
			ht[c]++
		} else {
			ht[c] = 1
		}
	}

	for _, c := range s2 {
		if count, ok := ht[c]; ok {
			if count > 1 {
				ht[c]--
			} else {
				delete(ht, c)
			}
		} else {
			return false
		}
	}

	if len(ht) > 0 {
		return false
	}

	return true
}
