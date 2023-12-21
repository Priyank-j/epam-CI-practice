package buytwochocolates

import "testing"

func TestBuyChoco(t *testing.T) {

	got := buyChoco([]int{1, 2, 3}, 3)
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
