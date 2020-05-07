package test

import (
	"testing"

	"../models"
)

func TestNewUser(t *testing.T) {
	user := models.NewUser("username", "password", "email")
	if user.Username != "username" || user.Password != "password" || user.Email != "email" {
		t.Error("No es posible crear el objeto", nil)
	}

}
