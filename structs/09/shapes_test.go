package structs

import (
	"testing"
)

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

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkArea(t, rectangle, 100.)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("Table driven tests", func(t *testing.T) {

		areaTests := []struct {
			shape Shape
			want  float64
		}{
			{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
			{shape: Circle{Radius: 10}, want: 314.1592653589793},
			{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
		}

		for _, tt := range areaTests {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		}

	})

}
