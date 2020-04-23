package orm

import (
	"time"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	CreatedAt time.Time
}

type Users []User

func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

func NewUser(username, password, email string) *User {
	user := &User { Username: username, Password: password, Email: email}
	return user
}

func GetUsers() Users {
	users := Users{}
	db.Find(&users)
	return users
}

func GetUser(id int) *User {
	user := &User{}
	db.Where("id=?", id).First(user)
	return user
}

func (this *User) Save() {
	db.Create(&this)

}