package main

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) string {
// 	// startA := time.Now()
// 	// // http.Get => returns an http.Response and an error
// 	// http.Get(a)
// 	// aDuration := time.Since(startA)

// 	// startB := time.Now()
// 	// http.Get(b)
// 	// bDuration := time.Since(startB)

// 	aDuration := measureResponseTime(a)
// 	bDuration := measureResponseTime(b)

// 	if aDuration < bDuration {
// 		return a
// 	}

// 	return b
// }

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

/************
use select for synchronizing process
************/

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	/*
		myVar := <-ch  (a blocking call. we need to wait for a value)
		select: wait on multiple channels
			here, which channel first have its code execuded in the select, which results its
			url being returned first
	*/
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// time.After will return a chan
	// prevent system blocking forever!
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// chan struct{}: here we don't care what type is sent to the channel
// struct{}: no memory allocation
func ping(url string) chan struct{} {
	// always "make" channels
	// if we declare it first (var ch chan struct{}), then it will be initialized with "zero" value of the type (nil)
	//    we cannot send to nil channels
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		// close a channel
		close(ch)
	}()
	return ch
}
