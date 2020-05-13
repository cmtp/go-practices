package test

import (
	"fmt"
	"math/rand"
	"testing"

	"../models"
)

const (
	username = "christian"
	password = "password"
	email    = "christian@example.com"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser(username, password, email)
	if user.Username != username || user.Password != password || user.Email != email {
		t.Error("No es posible crear el objeto", nil)
	}

}

func TestSave(t *testing.T) {
	user := models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(randomUsername(), password, email)
	if err != nil {
		t.Error("No es posible insertar el objeto", nil)
	}
}

func randomUsername() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}
