package validparenthesis

import (
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	got := isValid("{([])}()")
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
