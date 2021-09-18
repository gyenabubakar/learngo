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

		assertBalance(t, expected, actual)
	})

	t.Run("Withdraw", func(t *testing.T) {
		err := wallet.Withdraw(10)

		expected := Bitcoin(10)
		actual := wallet.Balance()

		assertNoError(t, err)
		assertBalance(t, expected, actual)
	})

	t.Run("Throw error on insufficient funds", func(t *testing.T) {
		// at this point, balance is 10 BTC, so withdrawing 20 BTC
		// should return an error
		err := wallet.Withdraw(20)

		assertError(t, err, InsufficientBalanceError)
	})
}

func assertBalance(tb testing.TB, expected, actual Bitcoin) {
	tb.Helper()

	if expected != actual {
		tb.Errorf("Expected %s. Received %s.", expected, actual)
	}
}

func assertError(tb testing.TB, err error, expectedErrorMessage string) {
	tb.Helper()

	if err == nil {
		tb.Fatal("Didn't get an expected error.")
	}

	if err.Error() != expectedErrorMessage {
		tb.Errorf("Expected %q. Received %q.\n", expectedErrorMessage, err.Error())
	}
}

func assertNoError(tb testing.TB, err error) {
	tb.Helper()

	if err != nil {
		tb.Errorf("There's an unchecked error.")
	}
}
