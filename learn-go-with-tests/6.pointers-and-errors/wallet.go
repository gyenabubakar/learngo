package banking

import "errors"

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(a Bitcoin) {
	w.balance += a
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(a Bitcoin) error {
	// if amount to withdraw is more than amount in wallet
	if a > w.balance {
		// return an error
		return errors.New("insufficient funds")
	}

	w.balance -= a
	return nil
}
