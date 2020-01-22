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
