package test

import (
	"fmt"
	"os"
	"testing"

	"../models"
)

func TestMain(m *testing.M) {
	beforeTest()
	result := m.Run()
	afterTest()
	os.Exit(result)
}

func beforeTest() {
	fmt.Println("Antes de las pruebas")
	models.CreateConnection()
	models.CreateTables()
}

func afterTest() {
	fmt.Println("Despues de las Pruebas Unitarias")
}
