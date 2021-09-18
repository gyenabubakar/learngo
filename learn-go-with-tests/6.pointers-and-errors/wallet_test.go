package banking

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	expected := Bitcoin(10)
	actual := wallet.Balance()

	if expected != actual {
		t.Errorf("Expected %s. Received %s.", expected, actual)
	}
}
