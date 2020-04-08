package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nombre := vars["nombre"]
	id := vars["id"]

	fmt.Fprintf(w, "Los parametros son "+nombre+" "+id)
	// w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/usuarios/{id:[0-9]+}", YourHandler).Methods("PUT", "DELETE")
	// DELETE, PUT, GET

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":3000", r))
}
