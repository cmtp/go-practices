package main

import (
	"fmt"
	"./models"
)

func main() {
	models.CreateConnection()
	models.CreateTables()

	user := models.NewUser("christian", "123")
	user.Save()
	fmt.Println(user)
	models.CloseConnection()
}
