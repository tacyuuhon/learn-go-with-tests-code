package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Panda!")
	want := "Hello, Panda!"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
