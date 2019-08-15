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

	t.Run("Area func", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := Area(rectangle)
		want := 100.

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Area()
		want := 100.

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

}
