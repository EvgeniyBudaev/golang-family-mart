package models

import "github.com/jmoiron/sqlx"

type User struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"password"`
}

func GetUserByUsername(Database *sqlx.DB, username string) (User, error) {
	user := User{}
	err := Database.Get(user, "SELECT EXISTS(id) FROM users WHERE username=?", username)
	return user, err
}
