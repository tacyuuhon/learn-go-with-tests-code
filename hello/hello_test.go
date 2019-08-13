package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	testHello(t, "Panda!")
}

func testHello(t *testing.T, name string) {
	got := hello(name)
	want := "Hello, " + name

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
