package iteration

import "testing"

func TestRepeat(t *testing.T) {

	t.Run("for loop style 01", func(t *testing.T) {
		repeated := Repeat01("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("for loop style 02", func(t *testing.T) {
		repeated := Repeat02("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})
}
