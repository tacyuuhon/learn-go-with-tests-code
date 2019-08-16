package pointers

import "fmt"

// Bitcoin type
type Bitcoin int

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit func.
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

// Balance func
func (w *Wallet) Balance() Bitcoin {
	fmt.Println("address of balance in Balance is", &w.balance)
	return w.balance
}
