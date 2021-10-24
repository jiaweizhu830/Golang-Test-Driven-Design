package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

// in order to mock Sleep
type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

//////////////////////////////////////////////////////////
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	/*
		Q: Why pass pointer here?
		A: It's pointer which implements the interface!!!
			In overriden Sleep() method, the receiver is pointer: *DefaultSleeper
	*/
	// sleeper := &DefaultSleeper{}
	// Countdown(os.Stdout, sleeper)

	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

/*
TDD: Take a thin slice of functionality and make it work end-to-end, backed by tests.

Note on mocking and spying
1. if more than 3 mocks then rethink on the design
	modularize the code
2. use spies with caution. Be sure you actually care about the details since it brings tighter
coupling between the test code and the implementation

Test Double:
 a generic term for any case where you replace a production object for testing purposes.
*/
