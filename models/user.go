package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/muneefdev/events-app/db"
	"github.com/muneefdev/events-app/utails"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
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

func (u *User) Validate() map[string]string {
	validate := validator.New()

	errors := map[string]string{}
	if err := validate.Struct(u); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			e.Tag()
			switch e.Tag() {
			case "required":
				errors[e.Field()] = " is required."
			case "email":
				errors[e.Field()] = "must be a valid email address."
			case "min":
				errors[e.Field()] = " must be at least " + e.Param() + " characters."
			case "max":
				errors[e.Field()] = "  must be at most " + e.Param() + " characters."
			default:
				errors[e.Field()] = " is invalid."
			}
		}
	}

	return errors
}
