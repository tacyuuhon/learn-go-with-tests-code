package structs

import "testing"

func TestPremeter(t *testing.T) {
	got := Premeter(0, 0)
	want := 0.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
