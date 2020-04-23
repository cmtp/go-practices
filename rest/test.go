package main

import (
	"fmt"
	"./models"
)

func main() {
	models.CreateConnection()
	models.CreateTables()

	user := models.CreateUser("christian", "123", "christian@admin.com")
	fmt.Println(user)
	models.CloseConnection()
}
