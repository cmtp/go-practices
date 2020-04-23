package models

// import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const userSchema string = `CREATE TABLE users(
		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(30) NOT NULL,
		password VARCHAR(64) NOT NULL,
		email VARCHAR(40),
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

type Users []User

func NewUser(username, password string) *User {
	user := &User { Username: username, Password: password}
	return user
}

func (this *User) Save() {
	sql := "INSERT users SET username=?, password=?"
	Exec(sql, this.Username, this.Password)
}
