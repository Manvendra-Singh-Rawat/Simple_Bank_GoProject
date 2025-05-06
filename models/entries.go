package models

import (
	"SimpleBank/db"
	"context"
)

func insertIntoEntries(accountID int, amount int) error {
	query := `INSERT INTO entries (account_id, amount) VALUES ($1, $2) RETURNING id`
	var fromEntriesID int
	err := db.Conn.QueryRow(context.Background(), query, accountID, amount).Scan(&fromEntriesID)
	return err
}
