package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
