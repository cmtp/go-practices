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

	user.Username = "Cambio de nombre"
	user.Password = "Cambio de Password"
	user.Email = "Cambio de Email"
	user.Save()
	models.CloseConnection()
}
