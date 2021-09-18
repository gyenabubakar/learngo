package banking

import "errors"

type Wallet struct {
	balance Bitcoin
}

var InsufficientBalanceError = "insufficient funds"

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
		return errors.New(InsufficientBalanceError)
	}

	w.balance -= a
	return nil
}
