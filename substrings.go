package strutil

// SubStrings returns all the substrings of a string
func SubStrings(s string) []string {
	substrings := make([]string, 0)
	if len(s) > 1 {
		for slen := 1; slen < len(s); slen++ {
			for i := 0; i+slen <= len(s); i++ {
				substrings = append(substrings, s[i:i+slen])
			}
		}
	}
	return substrings
}

// SubStringsMap returns all the substrings of a string in a map where:
// key = length of substring and value = slice of substrings of that (key) length
func SubStringsMap(s string) map[int][]string {
	sMap := make(map[int][]string)
	if len(s) > 1 {
		for slen := 1; slen < len(s); slen++ {
			substrings := make([]string, 0)
			for i := 0; i+slen <= len(s); i++ {
				substrings = append(substrings, s[i:i+slen])
			}
			if len(substrings) > 0 {
				sMap[slen] = substrings
			}
		}
	}
	return sMap
}
