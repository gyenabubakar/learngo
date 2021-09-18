package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(tb testing.TB, expected, actual Bitcoin) {
		tb.Helper()

		if expected != actual {
			tb.Errorf("Expected %s. Received %s.", expected, actual)
		}
	}

	wallet := Wallet{}

	t.Run("Deposit", func(t *testing.T) {
		wallet.Deposit(20)

		expected := Bitcoin(20)
		actual := wallet.Balance()

		assert(t, expected, actual)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet.Withdraw(10)

		expected := Bitcoin(10)
		actual := wallet.Balance()

		assert(t, expected, actual)
	})
}
