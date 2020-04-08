package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	UserName string
	Age      int
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseGlob("templates/**/*.html")
		if err != nil {
			panic(err)
		}
		user := User{"Christian", 22}
		template.Execute(w, user)
	})

	fmt.Println("El servidor a la escucha en el puerto :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
