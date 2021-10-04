package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// type StubStore struct {
// 	response string
// }

// func (s *StubStore) Fetch() string {
// 	return s.response
// }

/* To implement
1. takiing time to return data
2. a way to know it has been told to cancel
*/
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}

func TestServer(t *testing.T) {
	// data := "hello, world"
	// // http handler func
	// svr := Server(&StubStore{data})

	// request := httptest.NewRequest(http.MethodGet, "/", nil)
	// response := httptest.NewRecorder()

	// svr.ServeHTTP(response, request)

	// if response.Body.String() != data {
	// 	t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	// }

	t.Run("tell store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
		// if !store.cancelled {
		// 	t.Error("store was not told to cancel")
		// }

	})

	t.Run("return data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		store.assertWasNotCancelled()
		// if store.cancelled {
		// 	t.Error("it should not have cancelled the store")
		// }

	})
}

/*
context => help manage long-running process
	context package provides fns to derive new Context values from the existing ones
	When a context is cancelled, all contexts derived from it are also cancelled

	Incoming requests to a server should create a Context
	outgoing calls to servers should accept a Context
*/
