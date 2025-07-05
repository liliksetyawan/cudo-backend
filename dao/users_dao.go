package dao

import (
	"cudo/test/config"
	"cudo/test/repository"
)

func CreateUser(user repository.Users) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := config.DB.QueryRow(query, user.Name.String, user.Email.String).Scan(&user.ID)
	return err
}
