package main

import "testing"

func TestHello(t *testing.T) {
	// declaring and re-use values
	got := Hello("world")
	want := "Hello, world"

	if got != want {
		// print out a message and fail the test
		// f stands for format => so that we can use %q (placeholder)
		t.Errorf("got %q want %q", got, want)
	}
 }

 func TestHelloGreet(t *testing.T) {
	// declaring and re-use values
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		// print out a message and fail the test
		// f stands for format => so that we can use %q (placeholder)
		t.Errorf("got %q want %q", got, want)
	}
 }


 // Go is statically typed