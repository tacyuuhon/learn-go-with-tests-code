package pointers

import (
	"fmt"
	"testing"
)

// This is a wrong use case.
// In Go, when you call a function or a method the arguments are copied.
func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance() // It should return 10
	want := 0

	fmt.Println("address of balance in test is", &wallet.balance)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
