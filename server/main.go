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

func main() {
	http.HandleFunc("/", Hola) // DefaultServerMux

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: nil,
	}

	log.Fatal(server.ListenAndServe())
}
