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
	fmt.Println("ToString of balance in Deposit is", &w.balance)
	w.balance += amount
}

// Balance func
func (w *Wallet) Balance() Bitcoin {
	fmt.Println("ToString of balance in Balance is", &w.balance)
	return w.balance
}

// Withdraw func
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Stirng func over fmt.Stringer method
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
