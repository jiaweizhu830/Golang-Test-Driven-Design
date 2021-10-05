package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

// need to implement Handler interface to make a server!

/* type Handler interface {
	// arguments:
	// where we write our response
	// HTTP request that was sent to the server
	ServeHTTP(http.ResponseWriter, *http.Request)
 }
*/

/*
   // start a web server listening on a port on Handler
   // creating a goroutine for every request & running it against a Handler
   func ListenAndServe(addr string, handler Handler) error
*/

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

// Server
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWin(w, player)
	}
}

// GET /players/{name}
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	// if player not found, return 404 status code
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	// bc. both http.ResponseWriter and fmt.Fprint implement io Writer
	// fmt.Fprint(w, "20")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

// POST /players/{name}
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// Store
type InMemoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[name]++
}

// New Store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	// no need to initialize mutex
	return &InMemoryPlayerStore{store: map[string]int{}}
}

func main() {
	// http.HandlerFunc turns ordinary function into handler
	// handler := http.HandlerFunc(PlayerServer)

	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}

/*
go build: build program
./myProgram: execute

1. define interface
2. if only need to override a method defined in interface => use fn
3. if need to override a method defined in interface & has its own state => use struct
*/
