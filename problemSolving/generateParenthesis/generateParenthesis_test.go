package generateParenthesis

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	got := generateParenthesis(3)
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
