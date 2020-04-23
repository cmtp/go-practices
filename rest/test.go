package main

import (
	"fmt"
	"./models"
)

func main() {
	models.CreateConnection()
	models.CreateTables()

	models.CreateUser("christian1", "123", "christian@admin.com")
	models.CreateUser("christian2", "123", "christian@admin.com")
	models.CreateUser("christian3", "123", "christian@admin.com")
	
	users := models.GetUsers()
	fmt.Println(users)
	models.CloseConnection()
}
