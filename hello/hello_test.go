package main

import (
	"testing"
)

func TestHelloOk(t *testing.T) {
	testHello(t, "Hello, World!")
}

func TestHelloFail(t *testing.T) {
	testHello(t, "Welcome to Go world!")
}

func testHello(t *testing.T, want string) {
	got := hello()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
