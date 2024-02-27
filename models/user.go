package models

import (
	"errors"

	"github.com/muneefdev/events-app/db"
	"github.com/muneefdev/events-app/utails"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `binding:"required" json:"name"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	hashedPassword := utails.Hash(u.Password)
	u.Password = hashedPassword

	query := `
  INSERT INTO users (name, email, password)
  VALUES (?, ?, ?)`

	result, err := db.DB.Exec(query, u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	Id, _ := result.LastInsertId()
	u.ID = Id
	return err
}

func GetAllUsers() ([]User, error) {
	var users []User
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByEmail(userId string) (*User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, userId)
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (e *User) Login() (string, error) {
	user, err := GetUserByEmail(e.Email)
	if err != nil {
		return "", err
	}

	if !utails.ComparePasswords(user.Password, e.Password) {
		return "", errors.New("Invalid password")
	}

	token, err := utails.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", errors.New("Could not generate token")
	}

	return token, nil
}
