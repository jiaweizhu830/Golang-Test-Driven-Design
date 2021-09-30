package pointers

import (
	"errors"
	"fmt"
)

// this interface is defined in the fmt pacakge, and lets you define how the type is printed when
// using %s
/*
type Stringer interface {
	String() string
}
*/

// create type from existing ones: type MyName OriginalType
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// In Go, if starts with a lowercase symbol, then it's private outside the package it's defined in
	balance Bitcoin
}

// receiver variable: w (pointer to a wallet)
func (w *Wallet) Balance() Bitcoin {
	// here we can use both w.balance or (*w).balance
	return w.balance
}

// receiver variable: w (pointer to a wallet)
// no return
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in test is %v \n", &w.balance)

	// here we can use both w.balance or (*w).balance
	w.balance += amount
}

// var: define values global to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		// create a new error with a message
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

/**
In Go, when you call a function or a method, the arguments (e.g. receiver variable) are COPIED
In Go, permits to not explicit dereference a pointer (struct pointers: automatically dereferenced)
*/
