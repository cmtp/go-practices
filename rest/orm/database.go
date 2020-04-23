package orm

import (
	"../config"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// <username>:<password>@tcp(<host>:<port>)/<database>
func CreateConnection() {
	url := config.GetUrlDatabase()
	if connection, err := gorm.Open("mysql", url); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

func CloseConnection() {
	db.Close()
}

func CreateTables() {
	db.DropTableIfExists(&User{}) // Truncate
	db.CreateTable(&User{})
}