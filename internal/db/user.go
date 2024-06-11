package db

import (
	"database/sql"
	"errors"
	"homework/internal/models"
)


func GetUser(id int) (models.User, error) {
	var user models.User
	err := DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	result, err := DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)
	return nil
}

func UpdateUser(user models.User) error {
	_, err := DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	return err
}

func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}