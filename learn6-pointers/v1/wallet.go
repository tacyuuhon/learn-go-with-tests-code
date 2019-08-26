package pointers

import "fmt"

// Wallet struct
type Wallet struct {
	balance int
}

// Deposit func.
func (w Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

// Balance func
func (w Wallet) Balance() int {
	fmt.Println("address of balance in Balance is", &w.balance)
	return w.balance
}
