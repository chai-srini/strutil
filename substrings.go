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
