package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("HOST", "localhost") // create
	os.Unsetenv("HOST") // delete
	env := os.Getenv("HOST") // get
	fmt.Println(env)
}
