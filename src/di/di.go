package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Use interface io.Writer (now we can accept os.Stdout / bytes.Buffer / http.ResponseWriter)
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// func main() {
// 	Greet(os.Stdout, "Elodie")
// }

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
