package strutil

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		arg1 string
		arg2 string
		want bool
	}{
		{"Basic anagram 1", "ab", "ba", true},
		{"Basic anagram 2", "aabb", "baab", true},
		{"Basic non anagram", "ab", "ca", false},
		{"Basic single character anagram", "a", "a", true},
		{"Basic single character non anagram", "a", "b", false},
		{"Unequal parameters 1", "aa", "a", false},
		{"Unequal parameters 2", "", "a", false},
		{"Unequal parameters 3", "adsfasdf", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.arg1, tt.arg2); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
