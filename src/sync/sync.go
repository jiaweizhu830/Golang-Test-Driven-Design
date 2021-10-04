package sync

import "sync"

// Add lock (Mutext) to Counter

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()         // current goroutine acquires the lock
	defer c.mu.Unlock() // release the lock
	c.count++
}

func (c *Counter) Value() int {
	return c.count
}

/*
go vet // examine Go source code and report suspicious constructs
	Don't use embedding mutex!
		Readers may get access to Lock/Unlock from your public API
			=> coupling their own code

Mutex
	allow us to add locks to data
A Mutex must NOT be copied after first use
	pass Counter pointer, instead of creating a copy of Counter, then pass to other fns

WaitGroup
	means of waiting for goroutines to finish jobs


When to use locks over channels and goroutines?
	Use channels when passing ownership of data!!
	Use mutexes for managing state!!
*/
