package main

import (
	"fmt"
	"log"
	"net/http"

	"./mux"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world desde Hola")
}

type User struct {
	name string
}

func (this *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello User: "+this.name)
}

func main() {

	user := &User{"Christian"}

	mux := mux.CreateMux()
	mux.AddFunc("/hello", Hola)
	mux.AddHandle("/user", user)

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}
