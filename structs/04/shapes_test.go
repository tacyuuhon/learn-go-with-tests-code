package structs

import "testing"

func TestPremeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Premeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Area(rectangle)
	want := 100.

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
