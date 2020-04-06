package main

import (
	"fmt"
	"net/url"
)

func createURL() string {
	u, err := url.Parse("http://localhost:3000/params?nombre=valor")
	if err != nil {
		panic(err)
	}
	// u.Host = "localhost:3000"
	// u.Scheme = "http"

	// query := u.Query()
	// query.Add("nombre", "valor")

	// u.RawQuery = query.Encode()

	return u.String()
}

func main() {
	url := createURL()
	fmt.Println("La url final es: " + url)
}
