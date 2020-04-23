package main

import (
	"./orm"
)

func main() {
	orm.CreateConnection()
	orm.CloseConnection()
}
