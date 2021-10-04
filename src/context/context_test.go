package context

import (
	"context"
	"errors"
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
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	// start a slow process => append string to result
	go func() {
		var result string
		for _, c := range s.response {
			// goroutines listen to the ctx.Done(), will stop once a signal is sent in that channel
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	// wait for the goroutine to finish its work or for the cancellation to occur
	// make 2 async processes race each other to determine what we return
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}
// }

// func (s *SpyStore) assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}
// }

// in order to test we do NOT write anything to response on the error case
// SpyResponseWriter implements http.ResponseWriter
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
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

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
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
	})
}

/*
context => help manage long-running process
	context package provides fns to derive new Context values from the existing ones
	When a context is cancelled, all contexts derived from it are also cancelled

	Incoming requests to a server should create a Context
	outgoing calls to servers should accept a Context
*/
