package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(10)

	fmt.Println("ToString of balance in test is", &wallet.balance)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
