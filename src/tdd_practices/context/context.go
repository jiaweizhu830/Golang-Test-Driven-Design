package context

import (
	"context"
	"fmt"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Server code responsible for cancellation
		// ctx := r.Context()

		// data := make(chan string, 1)

		// go func() {
		// 	data <- store.Fetch()
		// }()

		// select {
		// case d := <-data:
		// 	// write data fetched from store to response
		// 	fmt.Fprint(w, d)
		// // Done() method returns a channel
		// case <-ctx.Done():
		// 	store.Cancel()
		// }

		// Server code NOT responsible for cancellation
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error)
	// Cancel()
}
