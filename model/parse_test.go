package model

import (
	"testing"
)

func TestTrimAll(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		// space
		{[]string{" ", " a ", "  a "}, []string{"", "a", "a"}},
		// magicSpace
		{[]string{magicSpace + "a" + magicSpace}, []string{"a"}},
	}
	for i, tt := range tests {
		trimAll(tt.input)
		if !eq(tt.input, tt.want) {
			t.Errorf("%v. got %v, want %v", i, tt.input, tt.want)
		}
	}
}

func TestReplaceSpace(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"  a  b  c  ", "a b c"},
		{" a      b c        ", "a b c"},
	}
	for i, tt := range tests {
		got := replaceSpace(tt.input)
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

func eq(s1, s2 []string) bool {
	if (len(s1) != len(s2)) || (cap(s1) != cap(s2)) {
		return false
	}
	for i, s := range s1 {
		if s != s2[i] {
			return false
		}
	}
	return true
}
