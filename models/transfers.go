package models

import (
	"SimpleBank/db"
	"context"
	"errors"
	"fmt"
)

type TransferMoney struct {
	FromAccountID int    `json:"FromAccountID" binding:"required"`
	ToAccountID   int    `json:"ToAccountID" binding:"required"`
	Amount        int    `json:"Amount" binding:"required"`
	Currency      string `json:"Currency" binding:"required,currency"`
}

func (transfer TransferMoney) CreateTransfer() error {
	query := "SELECT balance FROM accounts WHERE id = $1"
	var sendersAmount int
	err := db.Conn.QueryRow(context.Background(), query, transfer.FromAccountID).Scan(&sendersAmount)
	if err != nil {
		return err
	}

	if sendersAmount < transfer.Amount {
		return errors.New("sending amount exceeds amount in account")
	}

	query = "SELECT TransferMoney_SP($1, $2, $3)"
	_, err = db.Conn.Exec(context.Background(), query, transfer.FromAccountID, transfer.ToAccountID, transfer.Amount)
	if err != nil {
		fmt.Println("Error executing stored procedure: ", err)
		return err
	}
	return nil
}

func ValidAccount(accountID int64, currency string) error {
	accountDetails, err := GetAccount(accountID)
	if err != nil {
		return err
	}

	if accountDetails.Currency != currency {
		return errors.New("provided currency didn't match the currency in Database")
	}

	return nil
}
