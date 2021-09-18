package banking

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(a Bitcoin) {
	w.balance += a
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(a Bitcoin) {
	w.balance -= a
}
