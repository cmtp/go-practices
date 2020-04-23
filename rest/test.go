package main

import (
	"fmt"
	"./orm"
)

func main() {
	orm.CreateConnection()
	orm.CreateTables()

	orm.CreateUser("christian", "123", "christian@admin.com")
	orm.CreateUser("christian2", "123", "christian@admin.com")
	orm.CreateUser("christian3", "123", "christian@admin.com")
	
	// users := orm.GetUsers()
	// fmt.Println(users)

	user := orm.GetUser(1)
	user.Username = "testing"
	user.Password = "Cambio de Password"
	user.Email = "cambio de correo"
	user.Save()
	fmt.Println(user)
	orm.CloseConnection()
}
