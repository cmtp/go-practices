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

	user := orm.GetUser(1)
	user.Username = "testing"
	user.Password = "Cambio de Password"
	user.Email = "cambio de correo"
	user.Save()
	fmt.Println(user)
	user.Delete()
	orm.CloseConnection()
}
