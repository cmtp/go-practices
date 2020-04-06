package main

import (
	"fmt"
	"log"
	"net/http"
)

// handler
func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Hola2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World 2")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/dos", Hola2)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
