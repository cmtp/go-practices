package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "valor del header")
		w.Header().Add("NombreN", "valor del header")
		fmt.Fprintf(w, "Hola Mundo")
	})

	http.HandleFunc("/dos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo 2")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}