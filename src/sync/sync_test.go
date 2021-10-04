package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		/*
			sync.WaitGroup helps synchronize concurrent process

			A WaitGroup waits for a collection of goroutines to finish
				The main goroutine calls "Add" to set the #of goroutines to wait for
				Each of goroutines runs and calls "Done" when finished
				"Wait" can be used to block until all goroutines have finished
		*/
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		// Cannot copy Counter with Mutex => pass Counter pointer
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	// mark function as test helper
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

// return a Counter pointer
func NewCounter() *Counter {
	return &Counter{}
}
