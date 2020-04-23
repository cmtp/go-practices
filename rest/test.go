package main

import (
	"fmt"
	"./config"
)

func main() {
	url := config.GetUrlDatabase()
	fmt.Println(url)
}
