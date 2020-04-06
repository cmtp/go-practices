package main

import (
	"fmt"
	"log"
	"net/http"
)

// GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.URL)
		values := r.URL.Query()
		// remove query param
		values.Del("otro")
		values.Add("name", "Christian")
		values.Add("course", "Go Web")
		values.Add("Job", "Test")

		r.URL.RawQuery = values.Encode()

		fmt.Println(r.URL)
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
