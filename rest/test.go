package main

import (
	"fmt"

	"./models"
)

func main() {
	models.CreateConnection()
	models.Ping()
	result := models.ExistsTable("users")
	fmt.Println(result)
	models.CloseConnection()
}
