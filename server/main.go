package main

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World: "+this.name)
}

func main() {
	christian := &User{name: "Christian"}
	mux := http.NewServeMux()

	mux.Handle("/christian", christian)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
