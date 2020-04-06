package main

import (
	"fmt"
	"log"
	"net/http"
)

// GET, POST, PUT, DELETE

func main() {
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf(r.Header)

		accessToken := r.Header.Get("access_token")
		if len(accessToken) != 0 {
			fmt.Println(accessToken)
		}
		r.Header.Set("nombre", "valor")
		fmt.Println(r.Header.Get("nombre"))
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
