package main

import (
	"log"
	"net/http"
)

// func Redirect(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently)
// }

// func NotFound(w http.ResponseWriter, r *http.Request) {
// 	http.NotFound(w, r)
// }

func main() {

	redirect := http.RedirectHandler("https://www.google.com", http.StatusMovedPermanently)
	notFound := http.NotFoundHandler()

	mux := http.NewServeMux()

	mux.Handle("/redirect", redirect)
	mux.Handle("/not", notFound)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
