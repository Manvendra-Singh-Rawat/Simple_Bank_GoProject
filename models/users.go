package models

import (
	"SimpleBank/db"
	"context"
	"time"
)

type Users struct {
	UserName          string    `json:"UserName" binding:"required"`
	HashedPassword    string    `json:"HashedPassword" binding:"required"`
	FullName          string    `json:"FullName" binding:"required"`
	Email             string    `json:"Email" binding:"required"`
	PasswordChangedAt time.Time `json:"PasswordChangedAt"`
}

func (user *Users) CreateUser() error {
	query := `
		INSERT INTO users (username, hashed_password, full_name, email)
		VALUES ($1, $2, $3, $4) RETURNING username
	`
	var queryID string
	err := db.Conn.QueryRow(context.Background(), query, user.UserName, user.HashedPassword, user.FullName, user.Email).Scan(&queryID)
	if err != nil {
		return err
	}

	return nil
}
