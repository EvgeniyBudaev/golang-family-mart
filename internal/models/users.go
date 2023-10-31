package models

import (
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type User struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

func GetUserByUsername(Database *sqlx.DB, username string) (User, error) {
	user := User{}
	username = strings.ToLower(username)
	err := Database.Get(&user, "SELECT * FROM users WHERE username=$1", username)
	return user, err
}

func GetUserByEmail(Database *sqlx.DB, email string) (User, error) {
	user := User{}
	err := Database.Get(&user, "SELECT * FROM users WHERE email=$1", email)
	return user, err
}
func hashPassword(password string) string {
	return password
}

func AddUser(Database *sqlx.DB, user User) (User, error) {
	user.Username = strings.ToLower(user.Username)
	tx := Database.MustBegin()
	query := `
        INSERT INTO users(username, password, email) 
        VALUES($1, $2, $3)
        RETURNING id
    `
	row := tx.QueryRow(query, user.Username, hashPassword(user.Password), user.Email)
	var uuid string
	err := row.Scan(&uuid)
	if err != nil {
		log.Fatalln(err)
	}
	tx.Commit()
	user.Id = uuid
	return user, err
}
