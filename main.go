package main

import (
	"fmt"
	"log"
	"net/http"
)

// GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("el metodo es " + r.Method)

		switch r.Method {
		case "GET":
			fmt.Fprint(w, "Hello World from GET")
		case "POST":
			fmt.Fprint(w, "Hello World from POST")
		case "PUT":
			fmt.Fprint(w, "Hello World from PUT")
		case "DELETE":
			fmt.Fprint(w, "Hello World from DELETE")
		default:
			http.Error(w, "Invalid Method", http.StatusBadRequest)
		}
		// fmt.Fprintf(w, "Hola Mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
