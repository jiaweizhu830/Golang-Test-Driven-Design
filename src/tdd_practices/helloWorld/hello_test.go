package main

import "testing"

func TestHello(t *testing.T) {

	// assign functions to variables
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// to tell the test suite that this method is a helper
		t.Helper()

		if got != want {
			// print out a message and fail the test
			// f stands for format => so that we can use %q (placeholder for string)
			t.Errorf("got %q want %q", got, want)
		}
	}

	// subtests
	t.Run("saying hello to people", func(t *testing.T) {
		// declaring and re-use values
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		// declaring and re-use values
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// Go is statically typed

// TDD
// 1. create test for new requirements
// 2. fail test
// 3. add code
// 4. pass test
// 5. refactor
// 6. pass test
