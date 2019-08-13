package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("in Default language", func(t *testing.T) {
		got := Hello("Panda", "")
		want := "Hello, Panda"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("熊猫", "Chinese")
		want := "你好，熊猫"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("パンダさん", "Japanese")
		want := "こんにちは、パンダさん"
		assertCorrectMessage(t, got, want)
	})
}
