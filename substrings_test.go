package strutil

import (
	"reflect"
	"testing"
)

func TestSubStrings(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []string
	}{
		// test cases.
		{"Basic 3 char str", "abc", []string{"a", "b", "c", "ab", "bc"}},
		{"Empty", "", []string{}},
		{"Single char", "a", []string{}},
		{"two chars", "ab", []string{"a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStrings(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubStringsMap(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want map[int][]string
	}{
		{"Basic 3 char str", "abc", map[int][]string{1: {"a", "b", "c"}, 2: {"ab", "bc"}}},
		{"Empty", "", map[int][]string{}},
		{"Single char", "a", map[int][]string{}},
		{"two chars", "ab", map[int][]string{1: {"a", "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStringsMap(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubStringsMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
