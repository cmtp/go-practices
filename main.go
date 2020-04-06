package main

import (
	"fmt"
	"log"
	"net/http"
)

// GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println(r.URL.RawQuery)
		// fmt.Println(r.URL.Query())
		name := r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println(name)
		}

		param := r.URL.Query().Get("param")
		if len(param) != 0 {
			fmt.Println(param)
		}
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
