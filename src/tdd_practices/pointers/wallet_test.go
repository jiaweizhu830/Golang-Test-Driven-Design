package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

		// & => access pointer
		// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		// when use %s => call String() method on the type
		t.Errorf("got %s want %s", got, want)
	}
}

// assertError := func(t testing.TB, err error) {
// 	t.Helper()

// 	// nil is synonymous with null from other programming languages
// 	// error can be nil bc. the error is an interface

// 	/*
// 		A func that takes arguments or return values that are "interfaces", they can be "nillable"!
// 		When accessing a value that is nil, will throw a "runtime panic"!
//      "panic" will stop the program
// 	*/
// 	if err == nil {
// 		t.Error("wanted an error but didn't get one")
// 	}
// }

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		// Fatal => will stop the test if it's called
		// this way, will prevent "panic" when calling Error() from nil in the code below
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
