package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	t.Run("Deposit", func(t *testing.T) {
		wallet.Deposit(20)

		expected := Bitcoin(20)
		actual := wallet.Balance()

		if expected != actual {
			t.Errorf("Expected %s. Received %s.", expected, actual)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet.Withdraw(10)

		expected := Bitcoin(10)
		actual := wallet.Balance()

		if expected != actual {
			t.Errorf("Expected %s. Received %s.\n", expected, actual)
		}
	})
}
